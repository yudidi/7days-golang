# [gob解码编码流程 和 核心元素](https://www.processon.com/view/link/6047349207912947636e32a9)
// sendMsg -> network -> revData
// src1 -> desc1, src2 -> desc2
```
enc := gob.NewEncoder(&network) // desc1 // 构造函数
err := enc.Encode(&sendMsg) // src1

dec := gob.NewDecoder(&network)  // src2
err := dec.Decode(&revData) // desc2
```

[Golang Gob编码（gob包的使用）](https://blog.csdn.net/weixin_42117918/article/details/105864520)

# Q: gob编码的作用
A: 把xxx编码为bytes,字节流,这样可以在网络上传输
为了让某个数据结构能够在网络上传输或能够保存至文件，它必须被编码然后再解码。
当然已经有许多可用的编码方式了，比如 JSON、XML、Google 的 protocol buffers 等等。
而现在又多了一种，由Go语言 encoding/gob 包提供的方式。

> [gob编码,详解](http://c.biancheng.net/view/4597.html)

# Q: 为什么需要把数据结构进行编码？
有个小问题想请教下。 就是为什么序列化协议要搞这么多呢？ 从结构体变为字节流，固定一种方式不就好了吗。
[wps:自己动手写rpc](https://www.kdocs.cn/l/seuYeLKiXaus?f=130)
[](https://www.cnblogs.com/yudidi/p/14507632.html)