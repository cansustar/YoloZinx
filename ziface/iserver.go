package ziface

// 定义一个服务器接口

/*
通过定义这个接口，你可以创建不同类型的服务器，并确保它们都实现了相同的基本功能，例如启动、停止和运行。
同时，通过 AddRouter 方法，你可以为不同的服务器类型添加不同的路由处理逻辑。
*/

type IServer interface {
	Start() // 启动服务器
	Stop()  // 停止服务器
	Serve() // 运行服务器

	// AddRouter 路由功能： 给当前的服务注册一个路由方法，供客户端的链接处理使用
	AddRouter(router IRouter)
}

// 将在实现层模块znet中，实例化该接口
