package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		coder := Coder{conn: conn}
		//coder.Write(&Data{ii: 1})
		fmt.Printf("coder read 读取到gob解码的结构体: %+v \n", coder.Read())
	}
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

//func (coder *Coder) Write(goData interface{}) {
//	sendData := coder.encode(goData)
//	connWrite(coder.conn, sendData)
//}

func (c *Coder) encode(sendData interface{}) []byte {
	// TODO
	var storeEncodeResult bytes.Buffer
	enc := gob.NewEncoder(&storeEncodeResult) // // 指定存储编码结果的模块
	err := enc.Encode(&sendData)              // 数据编码后写入storeEncodeResult
	if err != nil {
		fmt.Println("编码失败", err)
	}
	return storeEncodeResult.Bytes()
}

type Data struct {
	ii int
}

func (c *Coder) decode(data []byte) Data {
	// TODO
	var storeEncodeResult bytes.Buffer
	dec := gob.NewDecoder(&storeEncodeResult) // 指定读取模块,读取待编码数据
	d := Data{}
	err := dec.Decode(&d) //传递参数必须为 地址
	if err != nil {
		fmt.Println("gob.Decode err:", err)
	}
	return d
}

// 真正读取数据
func connRead(conn net.Conn) []byte {
	buf := make([]byte, 1024)     // Make a buffer to hold incoming Data.
	reqLen, err := conn.Read(buf) // Read the incoming connection into the buffer.
	fmt.Println("server received", reqLen, string(buf), err)
	return buf
}

//// 真正写入数据
//func connWrite(conn net.Conn, sendData []byte) (n int, err error) {
//	return conn.Write(sendData)
//}
