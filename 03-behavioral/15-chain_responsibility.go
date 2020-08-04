package behavioral

import "fmt"

//职责链模式用于分离不同职责，并且动态组合相关职责。请求从链中的一个对象传到下一个对象，直到请求被响应为止。通过这种方式去除对象之间的耦合

var _ Handler = (*ProjManager)(nil)
var _ Handler = (*DepManager)(nil)
var _ Handler = (*Boss)(nil)

const (
	chainProj = iota
	chainDep
	chainBoss
)

//Handler 抽象处理者角色，定义一个处理请求的接口，包含抽象处理方法和一个后继连接
type Handler interface {
	CanHandle(funding int) bool
	Handle(funding int)
}

// * 链对象包含当前职责对象以及下一个职责链。
// * 职责对象提供接口表示是否能处理对应请求。
// * 职责对象提供处理函数处理相关职责

//RequestChain 链对象
type RequestChain struct {
	Handler               //职责对象
	next    *RequestChain //下一个职责链
}

//NewRequestChain 获取链对象
func NewRequestChain(chainType int) *RequestChain {
	var handler Handler
	switch chainType {
	case chainProj:
		handler = &ProjManager{}
	case chainDep:
		handler = &DepManager{}
	case chainBoss:
		handler = &Boss{}
	default:
		panic("never reach here")
	}
	return &RequestChain{
		Handler: handler,
	}
}

//Handle 链对象处理流程
func (chain *RequestChain) Handle(funding int) {
	if chain.CanHandle(funding) {
		chain.Handler.Handle(funding)
	} else if chain.next != nil {
		chain.next.Handle(funding)
	} else {
		fmt.Println("next chain is nil")
	}
}

//SetNexChain 设置下一个职责链&返回本身职责链
func (chain *RequestChain) SetNexChain(next *RequestChain) *RequestChain {
	chain.next = next
	return chain
}

//ProjManager 具体处理者角色-项目经理
type ProjManager struct {
}

//CanHandle 处理funding 3000以内
func (*ProjManager) CanHandle(funding int) bool {
	return funding <= 3000
}

//Handle 项目经理处理
func (*ProjManager) Handle(funding int) {
	fmt.Printf("proj mgr agree to apply for funding:%v  \n", funding)
}

//DepManager 具体处理者角色-部门经理
type DepManager struct {
}

//CanHandle 处理funding 6000以内
func (d *DepManager) CanHandle(funding int) bool {
	return funding <= 6000
}

//Handle 部门经理处理
func (d *DepManager) Handle(funding int) {
	fmt.Printf("dev mgr agree to apply for funding:%v  \n", funding)
}

//Boss 具体处理者角色-老板
type Boss struct {
}

//CanHandle 都可以处理
func (b *Boss) CanHandle(funding int) bool {
	return true
}

//Handle 老板Handle
func (b *Boss) Handle(funding int) {
	fmt.Printf("boss agree to apply for funding:%v  \n", funding)
}
