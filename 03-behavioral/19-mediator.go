package behavioral

import "fmt"

//中介者模式用一个中介对象来封装一系列对象交互，将多对多关联(网状结构)转换成一对多，构成星状结构

var _ Mediator = (*mediator)(nil)
var _ Colleague = (*colleague)(nil)

const (
	colleagueTypeMarket = iota
	colleagueTypeTech
)

// Mediator 抽象中介者角色,它是中介者的接口，提供了同事对象注册与转发同事对象信息的抽象方法
type Mediator interface {
	Operation(colleagueType int, data string)
	Registered(colleagueType int, c Colleague)
}

// mediator 具体中介者
// 实现中介者接口，定义一个 List 来管理同事对象，协调各个同事角色之间的交互关系，因此它依赖于同事角色
type mediator struct {
	list map[int]Colleague
}

//Registered 实现抽象中介Registered接口
func (m *mediator) Registered(colleagueType int, c Colleague) {
	c.SetMediator(m)
	if m.list == nil {
		m.list = make(map[int]Colleague)
	}
	m.list[colleagueType] = c
}

//Operation 实现抽象中介Operation接口
func (m *mediator) Operation(colleagueType int, data string) {
	cobj := m.list[colleagueType]
	switch cobj.(type) {
	case *Market:
		cobj.(*Market).Send(data)
	case *Technical:
		cobj.(*Technical).Receive(data)
	default:
		panic("never reach here")
	}
}

// NewMediator 获得中介对象
func NewMediator() Mediator {
	return &mediator{
		list: make(map[int]Colleague),
	}
}

// Colleague 抽象同事类角色,定义同事类的接口，保存中介者对象，提供同事对象交互的抽象方法，实现所有相互影响的同事类的公共功能
type Colleague interface {
	ReceiveMsg(data string)
	SendMsg(data string)
	SetMediator(Mediator)
}

// colleague 具体同事类角色,是抽象同事类的实现者，当需要与其他同事对象交互时，由中介者对象负责后续的交互
type colleague struct {
	mediator Mediator
}

// SetMediator 设置中介对象
func (c *colleague) SetMediator(m Mediator) {
	c.mediator = m
}

// SendMsg ...
func (c *colleague) SendMsg(data string) {
	fmt.Println(data)
}

// ReceiveMsg ...
func (c *colleague) ReceiveMsg(data string) {
	fmt.Println(data)
}

//Market 市场部
type Market struct {
	Colleague
}

// Send 发送消息
func (m *Market) Send(data string) {
	m.Colleague.SendMsg(data)
}

// Receive 接受消息
func (m *Market) Receive(data string) {
	m.Colleague.ReceiveMsg(data)
}

//Technical 技术部
type Technical struct {
	Colleague
}

// Send 发送消息
func (t *Technical) Send(data string) {
	t.Colleague.SendMsg(data)
}

// Receive 接受消息
func (t *Technical) Receive(data string) {
	t.Colleague.ReceiveMsg(data)
}
