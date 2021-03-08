package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type MsgData struct {
	X, Y, Z int
	Name    string
}

var network bytes.Buffer //网络传递的数据载体
// 对编码器,where to send the data,编码器把数据写到的这里去
// 对解码器.source of the data,从这里读取待解码的数据
func main() {
	err := senMsg()
	if err != nil {
		fmt.Println("编码错误")
		return
	}
	err = revMsg()
	if err != nil {
		fmt.Println("解码错误")
		return
	}
}

func senMsg() error {
	fmt.Print("开始执行编码（发送端）")

	enc := gob.NewEncoder(&network)
	sendMsg := MsgData{3, 4, 5, "jiangzhou"}
	fmt.Printf("原始数据：%+v \n", sendMsg)
	err := enc.Encode(&sendMsg)
	fmt.Println("传递的编码数据为：", network) // ydd: 编码后写到了这里,可看到编码结果
	return err
}
func revMsg() error {
	var revData MsgData
	dec := gob.NewDecoder(&network)
	err := dec.Decode(&revData) //传递参数必须为 地址
	fmt.Printf("解码之后的数据为：%+v \n", revData)
	return err
}
