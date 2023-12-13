package ziface

import "net"

// 定义链接模块的抽象层

type Iconnection interface {
	// Start 启动链接  让当前链接准备开始工作
	Start()
	// Stop 停止链接  结束当前链接的工作
	Stop()
	// GetTCPConnection 获取当前链接所绑定的socket conn
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取当前链接模块的链接ID
	GetConnID() uint32
	// RemoteAddr 获取远程客户端的TCP状态 IP port
	RemoteAddr() net.Addr
	// Send 发送数据，将数据发送给远程的客户端
	Send(data []byte) error
}

// HandleFunc 定义一个处理链接业务的方法  TODO!! 我才这里是接口型函数
// 这是一个处理业务的方法，
// 要处理业务，就要有一个形参是有关TCP链接，这个链接是客户端的链接句柄，因为是服务端处理业务，要返回给客户端，
// 后两个参数是：所要处理的业务的数据， 数据的长度
type HandleFunc func(*net.TCPConn, []byte, int) error
