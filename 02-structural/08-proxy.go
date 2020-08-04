package structural

import "fmt"

//代理模式用于延迟处理操作或者在进行实际操作前后对真实对象进行其它处理

var _ Subject = (*RealSubject)(nil)
var _ Subject = (*ProxyObject)(nil)

//Subject 抽象主题类
//通过接口或抽象类声明真实主题和代理对象实现的业务方法
type Subject interface {
	Do()
}

//RealSubject 真实主题类
//实现了抽象主题中的具体业务
type RealSubject struct {
}

// Do do实现
func (*RealSubject) Do() {
	fmt.Println("do something")
}

//ProxyObject 代理类
//提供了与真实主题相同的接口，其内部含有对真实主题的引用
type ProxyObject struct {
	RS *RealSubject
}

// Do 访问、控制或扩展真实主题的功能
func (p *ProxyObject) Do() {
	fmt.Println("proxy object ...")
	p.RS.Do()
	fmt.Println("proxy object end")
}
