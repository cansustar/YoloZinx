package ziface

/*
路由的抽象接口
/路由里的数据都是IRequest请求的
*/

type IRouter interface {
	// 模板的设计模式

	// PreHandle 在处理conn业务之前的钩子方法hook
	PreHandle(request IRequest)
	// Handle 在处理conn业务的主方法hook
	Handle(request IRequest)
	// PostHandle 在处理conn业务之后的钩子方法hook
	PostHandle(request IRequest)
}
