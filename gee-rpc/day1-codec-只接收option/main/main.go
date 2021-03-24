package main

import (
	"encoding/json"
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
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption) // TODO 这句话执行后就已经写入给conn了吗? 试试服务端是否可以读取到数据
	time.Sleep(time.Second)
}
