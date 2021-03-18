package main

import (
	"fmt"
	"net"
)

func main() {

}

// 实现从网络连接中读写数据
// 1.{结构体}需要编码为{字节流}后写入网络
// 2.从网络中获取的{字节流}需要解码为{结构体}
type Coder struct {
	conn net.Conn
}

func (coder *Coder) Read() interface{} {
	data := connRead(coder.conn)
	return coder.decode(data)
}

func (coder *Coder) Write(goData interface{}) {
	sendData := coder.encode(goData)
	connWrite(coder.conn, sendData)
}

func (c *Coder) encode(sendData interface{}) []byte {
	// TODO
	return nil
}

func (c *Coder) decode(data []byte) interface{} {
	// TODO
	return nil
}

// 真正读取数据
func connRead(conn net.Conn) []byte {
	buf := make([]byte, 1024)     // Make a buffer to hold incoming data.
	reqLen, err := conn.Read(buf) // Read the incoming connection into the buffer.
	fmt.Println("server received", reqLen, string(buf), err)
	return buf
}

// 真正写入数据
func connWrite(conn net.Conn, sendData []byte) (n int, err error) {
	return conn.Write(sendData)
}
