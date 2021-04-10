package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodecType   Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   GobType,
}

// 目的: 验证option和header的分隔符
// 结论: [10]换行符分隔,json.Decode一直读取到[10]的位置。
func main() {
	//编码(&Option{
	//	MagicNumber: MagicNumber,
	//	CodecType:   "\n",
	//})
	buf1 := 编码(&Option{
		MagicNumber: MagicNumber,
		CodecType:   GobType,
	})
	解码(buf1)
	buf2 := 编码(&Option{
		MagicNumber: MagicNumber,
		CodecType:   JsonType,
	})
	解码(buf2)

	解码(&bytes.Buffer{})
}

func 编码(o *Option) *bytes.Buffer {
	var buf bytes.Buffer // 这里也可以是conn
	// encode and write options too buf
	_ = json.NewEncoder(&buf).Encode(o)
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
	fmt.Println(buf.String())
	return &buf
}

func 解码(buf *bytes.Buffer) {
	var opt Option
	if err := json.NewDecoder(buf).Decode(&opt); err != nil {
		log.Println("decode options error: ", err)
		return
	}
	fmt.Printf("解码%+v\n", opt)
}
