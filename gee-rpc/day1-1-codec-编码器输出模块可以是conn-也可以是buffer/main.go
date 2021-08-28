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

type Data struct {
	MagicNumber int
	CodecType   Type
}

// 目的: json.NewEncoder,编码器的输出模块可以是conn, 也可以是buffer.
// 结论:
func main() {
	buf1 := 编码(Data{
		MagicNumber: MagicNumber,
		CodecType:   GobType,
	})
	解码(buf1)
}

func 编码(o Data) *bytes.Buffer {
	var buf bytes.Buffer // 这里也可以是conn
	// encode and write options too buf
	_ = json.NewEncoder(&buf).Encode(o)
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
	fmt.Println(buf.String())
	return &buf
}

func 解码(buf *bytes.Buffer) {
	var opt Data
	if err := json.NewDecoder(buf).Decode(&opt); err != nil {
		log.Println("decode options error: ", err)
		return
	}
	fmt.Printf("解码%+v\n", opt)
}
