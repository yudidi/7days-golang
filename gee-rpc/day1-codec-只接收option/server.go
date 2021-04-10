// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package geerpc

import (
	"encoding/json"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int        // MagicNumber marks this's a geerpc request
	CodecType   codec.Type // client may choose different Codec to encode body
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

// Server represents an RPC Server.
type Server struct{}

// NewServer returns a new Server.
func NewServer() *Server {
	return &Server{}
}

// DefaultServer is the default instance of *Server.
var DefaultServer = NewServer()

// ServeConn runs the server on a single connection.
// ServeConn blocks, serving the connection until the client hangs up.
func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()
	// 直接读取
	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	fmt.Println("服务端接收的原始字节流1:", n, err)
	fmt.Println("服务端接收的原始字节流2:", buf, string(buf))
	// 读取前面的固定JSON编码。字节[10]之前的部分.
	var opt Option
	//| Option{MagicNumber: xxx, CodecType: xxx} | Header{ServiceMethod ...} | Body interface{} |
	//| <------      固定 JSON 编码      ------>  10 <-------   编码方式由 CodeType 决定   ------->|
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		fmt.Println("rpc server: options error: ", err)
		return
	}
	fmt.Println("get opt", opt)

	//buf := make([]byte, 128)
	//n, err := conn.Read(buf)
	fmt.Println("server read after", n, err, buf, string(buf))
}

// Accept accepts connections on the listener and serves requests
// for each incoming connection.
func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		go server.ServeConn(conn)
	}
}

// Accept accepts connections on the listener and serves requests
// for each incoming connection.
func Accept(lis net.Listener) { DefaultServer.Accept(lis) }
