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

func main() {
	//编码(&Option{
	//	MagicNumber: MagicNumber,
	//	CodecType:   "/n",
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
	var buf bytes.Buffer
	// send options
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
