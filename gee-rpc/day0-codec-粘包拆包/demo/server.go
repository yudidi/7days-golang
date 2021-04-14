package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":4044")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen to 4044")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		} else {
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("关闭")
	fmt.Println("新连接：", conn.RemoteAddr())

	result := bytes.NewBuffer(nil)
	var buf [1024]byte
	var count int
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF { // TODO 表示没有读取到数据，而不是代表错误。
				//fmt.Print("s read err:", err)
				count++
				continue
			} else {
				fmt.Println("s read err:", err)
				break
			}
		} else {
			fmt.Println("recv:", result.String())
		}
		result.Reset()
		//fmt.Println("read io.EOF count",count)
	}
}
