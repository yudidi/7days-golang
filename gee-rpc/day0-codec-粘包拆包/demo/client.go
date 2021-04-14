package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 应用层交给 TCP 协议的数据并不会以消息为单位向目的主机传输，这些数据在某些情况下会被组合成一个数据段发送给目标的主机。
	// https://draveness.me/whys-the-design-tcp-message-frame/
	data := []byte("[这里才是一个完整的数据包]")
	fmt.Println(len(data))
	conn, err := net.DialTimeout("tcp", "localhost:4044", time.Second*30)
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	for i := 0; i < 1000; i++ {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
