package znet

import (
	"Yolozinx/ziface"
	"fmt"
	"net"
)

/*
链接模块
*/

type Connection struct {
	// 当前链接的Socket TCP套接字
	Conn *net.TCPConn

	// 链接的ID
	ConnID uint32

	// 当前链接的状态
	isClosed bool

	// 当前链接所绑定的处理业务的方法API
	//handleAPI ziface.HandleFunc  TODO 已被Router取代

	// 告知当前链接已经退出/停止的channel
	ExitChan chan bool

	// 该链接处理的方法Router
	Router ziface.IRouter
}

// NewConnection 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:   conn,
		ConnID: connID,
		//handleAPI: callback_api,  TODO 已被Router替代
		Router:   router,
		isClosed: false, // 刚创建链接，当前链接为开启状态
		ExitChan: make(chan bool, 1),
	}
	return c
}

// StartReader 链接的读方法 Reader 从客户端去读数据，交给绑定的handle处理业务
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID=", c.ConnID, "Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	// 做一个循环从客户端读数据， 然后执行绑定的业务
	for true {
		// 读取客户端的数据到buf中，目前最大是512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf) // 把buf放到request里了，所以这里就用不到cnt了
		if err != nil {
			fmt.Println("recv buf err", err)
			// 读失败的话， continue
			continue
		}
		// TODO 这部分功能由路由代为实现
		//// 调用当前链接所绑定的HandleAPI
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("ConnID", c.ConnID, "handle is error", err)
		//	break
		//}
		//// 这里就进入了handleAPI的业务逻辑，所以后面没东西了

		// 得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf,
		}
		// 从路由中，找到注册绑定的Conn对应的router调用
		go func(request ziface.IRequest) {
			// 执行注册的路由方法
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}

}

// Start 启动链接  让当前链接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start()...ConnID=", c.ConnID)
	// 启动从当前链接的读数据的业务
	// start方法 开启一个Reader
	go c.StartReader()
	// TODO 启动从当前链接写数据的业务

}

// Stop 停止链接  结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID=", c.ConnID)

	// 如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// 关闭socket连接
	c.Conn.Close()

	// 关闭管道, 回收资源
	close(c.ExitChan)

}

// GetTCPConnection 获取当前链接所绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {

	return c.Conn

}

// GetConnID 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {

	return c.ConnID

}

// RemoteAddr 获取远程客户端的TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {

	return c.Conn.RemoteAddr()

}

// Send 发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
