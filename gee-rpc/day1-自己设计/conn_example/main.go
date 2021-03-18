package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

//https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go
// If you build this and run it, you'll have a simple TCP server running on port 3333.
// TODO 命令行作为client给tcp或udp端口写入数据
//  To test your server, send some raw data to that port:
//  echo -n "test out the server" | nc localhost 3333
func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf) // TODO 使用网络连接读取数据
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("server received: " + string(buf) + "\n")
	// TODO 使用网络连接输出数据
	// Send a response back to person contacting us.
	conn.Write([]byte(fmt.Sprintf("Message received: %v %v \n", reqLen, string(buf))))
	// Close the connection when you're done with it.
	conn.Close()
}
