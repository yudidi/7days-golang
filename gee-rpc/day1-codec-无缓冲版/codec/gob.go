package codec

import (
	"encoding/gob"
	"io"
	"log"
)

// 这个结构体由四部分构成，conn 是由构建函数传入，通常是通过 TCP 或者 Unix 建立 socket 时得到的链接实例，
// dec 和 enc 对应 gob 的 Decoder 和 Encoder，
// buf 是为了防止阻塞而创建的带缓冲的 Writer，一般这么做能提升性能。
type GobCodec struct {
	conn io.ReadWriteCloser
	dec  *gob.Decoder
	enc  *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

// 传入1个套接字连接(或者1个实现了3个接口的接口对象)
func NewGobCodec(connReaderWriter io.ReadWriteCloser) Codec {
	//bufWriter := bufio.NewWriter(connReader) // conn具备从网卡写入到内存的能力.
	return &GobCodec{
		conn: connReaderWriter,
		dec:  gob.NewDecoder(connReaderWriter), // 解码器构造时需要读取模块(具备Reader能力),读取模块不能改变.
		enc:  gob.NewEncoder(connReaderWriter), // 编码器构造时存储模块(具备存储能力),并且存储模块不能改变.
	}
}

func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h) // 解码到header
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body) // 解码到body
}

// 编码
func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		if err != nil {
			_ = c.Close()
		}
	}()
	// 把header的编码结果写入编码器的存储模块
	if err = c.enc.Encode(h); err != nil {
		log.Println("rpc: gob error encoding header:", err)
		return
	}
	// 把body的编码结果写入编码器的存储模块
	if err = c.enc.Encode(body); err != nil {
		log.Println("rpc: gob error encoding body:", err)
		return
	}
	return
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}
