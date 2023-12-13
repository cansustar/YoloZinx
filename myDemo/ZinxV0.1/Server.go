package main

import "Yolozinx/znet"

/*
基于Zinx框架开发的，服务器端应用程序
*/

func main() {
	// 创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.1]")
	// 启动server
	// 为什么用户通过Serve来启动，而不是Start呢？
	// 因为我们并不希望start和stop把这两个操作暴露给用户
	s.Serve()
}
