package behavioral

import "fmt"

//将请求封装成一个对象，从而可以用不同的请求对客户进行参数化，实现调用者和接收者的解藕

var _ Command = (*command)(nil)
var _ Receiver = (*receiver)(nil)

const (
	buttonOpen = iota
	buttonClose
)

//Command 抽象命令类角色，声明执行命令的接口
type Command interface {
	Execute()
}

//command 具体命令角色，抽象命令类的具体实现类，它拥有接收者对象，并通过调用接收者的功能来完成命令要执行的操作
type command struct {
	Receiver
	Button int
}

//Execute 执行命令接口实现
func (c command) Execute() {
	c.Action(c.Button)
}

//Receiver 接口
type Receiver interface {
	Action(button int)
}

//receiver 实现者/接收者角色
type receiver struct {
}

func (r receiver) Action(button int) {
	switch button {
	case buttonClose:
		fmt.Println("关闭电视机")
	case buttonOpen:
		fmt.Println("打开电视机")
	}
}

//Invoker 调用者/请求者角色,是请求的发送者，它通常拥有很多的命令对象，并通过访问命令对象来执行相关请求，它不直接访问接收者
type Invoker struct {
	list []Command
}

//AddCommand 添加命令
func (invoker *Invoker) AddCommand(list ...Command) {
	invoker.list = append(invoker.list, list...)
}

//ExecuteCommand 执行命令
func (invoker *Invoker) ExecuteCommand() {
	for _, c := range invoker.list {
		c.Execute()
	}
}
