package znet

import (
	"Yolozinx/ziface"
	"fmt"
	"net"
)

// Server 定义一个Server的服务器模块， IServer的接口实现
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
}

// Start 实现了IServer接口的Start方法 启动服务器
func (s *Server) Start() {
	// 开发一个单体服务器的步骤
	fmt.Printf("[Start]Server Listenner at IP: %s, Port:%d, is starting\n", s.IP, s.Port)

	go func() {
		// 1. 获取一个TCP的Addr(套接字)
		// addr就是监听的socket 句柄
		/*
			Socket句柄是一个用于标识和操作网络套接字的抽象概念。
			在网络编程中，套接字（Socket）是用于在网络上进行通信的一种机制，
			而套接字句柄（Socket Handle）则是对套接字的引用，允许程序对套接字进行读取、写入、关闭等操作。
		*/
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addt error:", err)
			return
		}
		//2. 尝试监听这个服务器地址
		// 成功的话，就拿到了监控的句柄
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start Zinx server success,", s.Name, "succ, Listenning....")
		// 监听本地端口已经成功
		//3. 阻塞等待客户端进行连接，处理客户端连接业务(读写) 这里会阻塞，所以改成异步的操作
		for {
			// 如果有客户端连接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// 已经与客户端建立连接，做一些业务， 做一个最基本的512字节长度的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					//将buf这个字节切片，通过conn的Read方法，读到这个buf中，返回一个成功读的字节和错误编码，cnt表示已经读到的字节大小
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						return
					}
					// 如果读成功
					fmt.Printf("recv client buf %s, cnt %d \n", buf, cnt)
					// 回显功能  将从开头， 到已经读到的位置写到buf里
					if _, err := conn.Write(buf[0:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()

}

// Stop 实现了IServer接口的Stop方法 停止服务器
func (s *Server) Stop() {
	// TODO 将一些服务器的资源，状态，或者一些已经开辟的连接信息，进行停止或者回收

}

// Serve 实现了IServer接口的Serve方法 运行服务器
func (s *Server) Serve() {
	// 客户端通过Serve方法来启动zinx,所以在Serve里应该调用Start
	s.Start()
	// Start本身是异步的，希望在Serve阻塞

	// TODO 做一些启动服务器之后的，额外业务工作

	// 阻塞状态
	select {}
}

// NewServer 初始化Server模块的方法  应该返回一个抽象层的IServer
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
