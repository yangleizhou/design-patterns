package behavioral

import "fmt"

//多个对象间存在一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。这种模式有时又称作发布-订阅模式、模型-视图模式

var _ Subject = (*subject)(nil)
var _ Observer = (*observer)(nil)
var _ Observer = (*InvestorObserver)(nil)

//Subject 抽象目标类，它提供了一个用于保存观察者对象的聚集类和增加、删除观察者对象的方法，以及通知所有观察者的抽象方法
type Subject interface {
	Register(Observer) Subject
	Remove(Observer)
	Notify(ctx string)
}

// subject 具体目标类，它实现抽象目标中的通知方法，当具体主题的内部状态发生改变时，通知所有注册过的观察者对象
type subject struct {
	list []Observer
	ctx  string
}

// Register 添加观察者对象
func (s *subject) Register(o Observer) Subject {
	s.list = append(s.list, o)
	o.SetRegisterIndex(len(s.list) - 1)
	return s
}

// Remove 移除观察者对象
func (s *subject) Remove(o Observer) {
	index := o.GetRegisterIndex()
	if index > len(s.list)-1 {
		return
	}
	if s.list[index] != o {
		return
	}
	s.list = append(s.list[:index], s.list[index+1:]...)
	for i := index; i < len(s.list); i++ {
		s.list[i].SetRegisterIndex(i)
	}
}

// Notify 通知
func (s *subject) Notify(ctx string) {
	s.ctx = ctx
	for _, o := range s.list {
		o.Update(s.ctx)
	}
}

//Observer 抽象观察者角色，它是一个抽象类或接口，它包含了一个更新自己的抽象方法，当接到具体目标的更改通知时被调用
type Observer interface {
	SetRegisterIndex(index int)
	GetRegisterIndex() int
	Update(ctx string)
}

// observer 具体观察者角色,实现抽象观察者中定义的抽象方法，以便在得到目标的更改通知时更新自身的状态
type observer struct {
	registerIndex int
	Name          string
}

// SetRegisterIndex 设置注册索引index
func (o *observer) SetRegisterIndex(index int) {
	fmt.Println("observer index = ", index)
	o.registerIndex = index
}

// GetRegisterIndex 获取注册的索引index
func (o *observer) GetRegisterIndex() int {
	return o.registerIndex
}

// Update 当接到具体目标subject的更改通知时被调用
func (o *observer) Update(ctx string) {

}

// InvestorObserver 股票投资者
type InvestorObserver struct {
	Name string
	*observer
}

//Update 实现观察者update
func (io *InvestorObserver) Update(ctx string) {
	fmt.Printf(" %s receive %s\n", io.Name, ctx)
}

//NewInvestorObserver 创建observer
func NewInvestorObserver(name string) Observer {
	return &InvestorObserver{
		Name:     name,
		observer: &observer{},
	}
}
