package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

var network bytes.Buffer //网络传递的数据载体

func main() {
	err := senMsg()
	if err != nil {
		fmt.Println("编码错误")
		return
	}
	err = revMsg()
	if err != nil {
		fmt.Println("解码错误", err)
		return
	}
}

// 1.自我描述的编码方式，所以自己流头部，说明了这个字节流的长度。那么解码时就知道读取多少给字节。
// 2.什么叫自我描述的编码方式
// A: 就是像json，xml一样,不需要定义proto或者thrift协议文件,这样更易用。不过就需要多浪费一些空间用来自我描述。
func senMsg() error {
	//mock := []byte{5, 12, 0, 2, 49, 48} // "10"编码结果
	mock := []byte{5, 12, 0, 2, 49, 48, 5, 6} // 在"10"编码结果后,插入2个字节
	network.Write(mock)
	fmt.Printf("写入后的缓存区 %+v \n", network)
	return nil
}

func revMsg() error {
	var revData string
	dec := gob.NewDecoder(&network) // 从network读取数据,然后解码 // 解码器需要1个读取数据的东西
	err := dec.Decode(&revData)     //传递参数必须为 地址
	fmt.Printf("解码之后的数据为：%+v \n", revData)
	fmt.Printf("读取后的缓存区 %+v", network)
	return err
}
