package structural

import "fmt"

//将抽象与实现分离，使它们可以独立变化。它是用组合关系代替继承关系来实现的，从而降低了抽象和实现这两个可变维度的耦合度

var _ AbstractMessage = (*RefinedMessage)(nil)
var _ MessageImplementer = (*MessageTCP)(nil)
var _ MessageImplementer = (*MessageHTTP)(nil)

// AbstractMessage 抽象化角色
type AbstractMessage interface {
	SetMessageImplementer(MessageImplementer)
	SendMessage(text, to string)
}

// RefinedMessage 扩展抽象化角色
// 抽象部分的具体实现，该类一般是对抽象部分的方法进行完善和扩展。
type RefinedMessage struct {
	mi MessageImplementer
}

// SendMessage 抽象化角色接口部分的具体实现
func (rm *RefinedMessage) SendMessage(text, to string) {
	rm.mi.Send(text, to)
}

// SetMessageImplementer 赋值实现化角色
func (rm *RefinedMessage) SetMessageImplementer(mi MessageImplementer) {
	rm.mi = mi
}

// NewRefindMessage 获取抽象角色对象
func NewRefindMessage() AbstractMessage {
	return &RefinedMessage{}
}

// MessageImplementer 实现化角色
//实现化角色,这个角色给出实现化角色的接口，但不给出具体的实现
//实现化接口与抽象化接口可以不一样,实现化角色应当只给出底层操作，而抽象化角色应当只给出基于底层操作的更高一层的操作
type MessageImplementer interface {
	Send(text, to string)
}

// MessageTCP 具体实现化角色
// 实现化角色接口的具体实现
type MessageTCP struct {
}

// Send 实现实现化接口
func (mt *MessageTCP) Send(text, to string) {
	fmt.Printf("send %s to %s by TCP\n", text, to)
}

// NewMessageTCP 获取具体对象
func NewMessageTCP() MessageImplementer {
	return &MessageTCP{}
}

// MessageHTTP 具体实现化角色
// 实现化角色接口的具体实现
type MessageHTTP struct {
}

// Send 实现实现化接口
func (mt *MessageHTTP) Send(text, to string) {
	fmt.Printf("send %s to %s by HTTP\n", text, to)
}

// NewMessageHTTP 获取具体对象
func NewMessageHTTP() MessageImplementer {
	return &MessageHTTP{}
}
