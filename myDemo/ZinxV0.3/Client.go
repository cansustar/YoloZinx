package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/

func main() {
	fmt.Println("client start...")
	// 1. 直接连接远程服务器，得到一个conn连接

	time.Sleep(time.Second)

	// 创建一个对话
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		// 2. 连接调用Write方法，写数据
		_, err := conn.Write([]byte("Hello Zinx V0.2.."))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}

		buf := make([]byte, 512)
		// 读取数据到buf中
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		// server返回了个数据
		fmt.Printf("server call back:%s, cnt=%d \n", buf, cnt)

		// cpu阻塞, 如果一直在上面判断的话可能会把cpu跑满
		// 每隔一秒循环一次
		time.Sleep(time.Second)
	}
}
