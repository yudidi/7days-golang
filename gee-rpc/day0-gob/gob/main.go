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

type MsgData2 struct {
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

// sendMsg -> network -> revData
// sendMsg -> network -> revData
// src1 -> desc1, src2 -> desc2 // TODO 解码和编码器需要desc1和src2是不变的,所以作为构造函数的参数。 而src1和desc2是变化的
// TODO 进一步可对照物理世界:
//  1个密码编码器创建时，存储模块(能够写入东西的模块Write)就一起创建了，但是需要编码的东西则还不确定(所以是interface{})
//  1个解码器创建时,读取模块(能够读取密文的模块Read)就一起创建了,但是解码的结果还不确定(所以是interface{})
func senMsg() error {
	fmt.Print("开始执行编码（发送端）")
	enc := gob.NewEncoder(&network) // 数据编码后写入network // 编码器需要1个存放编码结果的地方.
	sendMsg := MsgData{3, 4, 5, "jiangzhou"}
	fmt.Printf("原始数据：%+v \n", sendMsg)
	err := enc.Encode(&sendMsg)
	fmt.Println("传递的编码数据为：", network) // ydd: 编码后写到了这里,可看到编码结果
	return err
}

func revMsg() error {
	var revData MsgData
	dec := gob.NewDecoder(&network) // 从network读取数据,然后解码 // 解码器需要1个读取数据的东西
	err := dec.Decode(&revData)     //传递参数必须为 地址
	fmt.Printf("解码之后的数据为：%+v \n", revData)
	return err
}
