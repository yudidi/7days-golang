[Golang Gob编码（gob包的使用）](https://blog.csdn.net/weixin_42117918/article/details/105864520)

# Q: gob编码的作用
A: 把xxx编码为bytes,字节流,这样可以在网络上传输
为了让某个数据结构能够在网络上传输或能够保存至文件，它必须被编码然后再解码。
当然已经有许多可用的编码方式了，比如 JSON、XML、Google 的 protocol buffers 等等。
而现在又多了一种，由Go语言 encoding/gob 包提供的方式。

> [gob编码,详解](http://c.biancheng.net/view/4597.html)

# Q: 为什么需要把数据结构进行编码？
https://www.kdocs.cn/l/seuYeLKiXaus?f=130