package codec

import (
	"io"
)

// TODO 类似协议设计
// 客户端请求 = 服务名 + 方法名 + 参数 + 其他信息(header)
// 服务端响应 = error + reply
// reply 和 args 抽象为body

// 客户端发送的请求包括服务名 Arith，方法名 Multiply，参数 args 三个，服务端的响应包括错误 error，返回值 reply 2 个。
//  我们将请求和响应中的参数和返回值抽象为 body，剩余的信息放在header中
// TODO  请求和响应消息,共用的headr
type Header struct {
	ServiceMethod string // format "Service.Method"
	Seq           uint64 // sequence number chosen by client
	// 错误信息，客户端置为空，服务端如果如果发生错误，将错误信息置于 Error 中。
	Error string
}

//  TODO 客户端和服务端公用的一套编码解码协议
//  写入: 消息体和消息头
//  读取: 消息头
//  读取: 消息体
// 抽象出对消息体进行编解码的接口 Codec，抽象出接口是为了实现不同的 Codec 实例：
type Codec interface {
	// 接口嵌入. Codec定义Closer的接口方法
	io.Closer
	// 读取客户端或服务端 消息体的header
	ReadHeader(*Header) error
	// 读取客户端请求的args 或 服务端响应的reply, 统称body
	ReadBody(interface{}) error
	// 往header写入内容
	// 把header和body写入网络连接
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

// 这部分代码和工厂模式类似，与工厂模式不同的是，返回的是构造函数，而非实例。
var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
