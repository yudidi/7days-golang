package main

import (
	"encoding/json"
	"fmt"
	"geerpc"
	"log"
	"net"
	"time"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	// 模拟服务端处理
	geerpc.Accept(l)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	// in fact, following code is like a simple geerpc client
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()
	// send options
	bytes, err := json.Marshal(geerpc.DefaultOption)
	fmt.Println("客户端发送的数据", bytes, err)
	// TODO 查看json.Encode API,内部有一句"e.WriteByte('\n')",使得序列化之后比原始数据多了1个字节('\n')
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption) // 这句话执行后就已经写入给conn了,服务端可以获取到数据
	time.Sleep(2 * time.Second)
}
