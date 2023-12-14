package main

import (
	"Yolozinx/ziface"
	"Yolozinx/znet"
	"fmt"
)

/*
基于Zinx框架开发的 服务器端应用程序
*/

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

// Handle Test
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...\n"))
	if err != nil {
		fmt.Println("call back ping error")
	}
}

// PostHandle Test
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router afterHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}

func main1() {
	// 1. 创建一个server句柄， 使用Zinx的API
	s := znet.NewServer("[zinx V0.3]")

	// 给当前zinx框架添加一个自定义的router
	s.AddRouter(&PingRouter{})

	// 3. 启动server
	s.Serve()
}
