package main

import (
	"encoding/json"
	"fmt"
	"geerpc"
	"geerpc/codec"
	"log"
	"net"
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
	//time.Sleep(time.Second)
	// send options
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption) // TODO 这句话往conn里面写入了啥东西
	buf := make([]byte, 1024)
	fmt.Printf("conn %+v \n", conn)
	n, err := conn.Read(buf)
	fmt.Println("read result", n, err)
	fmt.Println("写入了 ", string(buf))
	cc := codec.NewGobCodec(conn)
	// 模拟客户端
	// send request & receive response
	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		// 模拟客户端端发送数据
		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
		// 模拟客户端接收数据 // 读取响应头
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("响应：", "h:", h, "reply:", reply)
	}
}
