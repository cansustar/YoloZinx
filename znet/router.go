package znet

import "Yolozinx/ziface"

// BaseRouter BaseRouter里不需要提供方法， 具体的实现应该是由用户来自定义的。
// 有两种方式，一种是用户去继承重写IRouter, 要么就先定义一个BaseRouter，先把这些功能实现，然后由用户来重写

// BaseRouter 实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类的方法进行重写
type BaseRouter struct {
}

// 这里之所以BaseRouter的方法都为空，是因为用户自己实现的Router不一定希望实现PreHandle，PostHandle这两个业务
// 所以用户自定义的Router继承BaseRouter而不是去实现Irouter的好处就是，不需要实现全部方法

// PreHandle 在处理conn业务之前的钩子方法Hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}

// Handle 在处理conn业务的主方法hook
func (br *BaseRouter) Handle(request ziface.IRequest) {

}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}
