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

// 传入1个套接字连接, 返回基于这个conn进行数据传输且使用gob序列化协议的 编码和解码器.
func NewGobCodec(conn io.ReadWriteCloser) Codec {
	// 默认:4096 // 这里只有Write方法是走缓存的 // TODO 可以了解bufio.Write方法的逻辑 //因为一般写入的内容不会超过buf,所以需要主动flush.
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		// 对结构体进行编码,经过缓存区后,输出到buf中 // TODO 结构体->序列化->buf()->buf内部的conn->网线 // c or s发送时走缓存,减少socket系统调用.
		enc: gob.NewEncoder(buf),
		// 从conn读取字节流,然后解码为结构体 // TODO ->网线->字节流->反射->结构体
		dec: gob.NewDecoder(conn),
	}
}

func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

// 因为是带缓冲的写,所以每次都需要flush buf to 底层的conn
func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	// 把header和body两个结构体序列化之后,统一写入到buf中 // TODO 考虑body过大，超过缓存区大小的情况，会导致部分body先发送另一端。
	if err = c.enc.Encode(h); err != nil {
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
