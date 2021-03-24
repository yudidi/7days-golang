package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

// 这个结构体由四部分构成，conn 是由构建函数传入，通常是通过 TCP 或者 Unix 建立 socket 时得到的链接实例，
// dec 和 enc 对应 gob 的 Decoder 和 Encoder，
// buf 是为了防止阻塞而创建的带缓冲的 Writer，一般这么做能提升性能。
type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer // TODO 客户端和服务端调用Write方法时，都是先写到buf(默认4096),然后再统一发送给对端。
	dec  *gob.Decoder
	enc  *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

// 传入1个套接字连接
func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn) // 默认:4096
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		// TODO 这个时候才从缓存区写入到底层Write(conn中). header和body一起发送下去
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err = c.enc.Encode(h); err != nil { //
		log.Println("rpc: gob error encoding header:", err)
		return
	}
	//c.buf.WriteString("11")
	if err = c.enc.Encode(body); err != nil {
		log.Println("rpc: gob error encoding body:", err)
		return
	}
	return
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}
