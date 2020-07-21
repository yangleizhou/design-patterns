package patterns

import "fmt"

//The intent of the Builder design pattern is to separate the construction of a complex object from its representation.
//By doing so the same construction process can create different representations.

//编译期将*MacBookBuilder 转换为Builder
//如果接口Builder发生变化，而MacBookBuilder没有变化，编译将不能通过
var _ Builder = (*MacBookBuilder)(nil)

// Builder 构建者的抽象基类
type Builder interface {
	SetType() Builder
	SetCPU() Builder
	SetOS() Builder
	GetProduct() ComputerProduct
}

// Director 负责组装方法
type Director struct {
	builder Builder
}

// MakeProduct 组装
func (d *Director) MakeProduct(builder Builder) {
	d.builder = builder
	d.builder.SetType().SetCPU().SetOS()
}

// ComputerProduct 最终对象
type ComputerProduct struct {
	Type string
	CPU  string
	OS   string
}

// MacBookBuilder 苹果笔记本
type MacBookBuilder struct {
	computer ComputerProduct
}

// SetType 电脑类型
func (mb *MacBookBuilder) SetType() Builder {
	mb.computer.Type = "top"
	return mb
}

// SetCPU 设置CPU
func (mb *MacBookBuilder) SetCPU() Builder {
	mb.computer.CPU = "intel i7"
	return mb
}

// SetOS 设置OS
func (mb *MacBookBuilder) SetOS() Builder {
	mb.computer.OS = "mac"
	return mb
}

// GetProduct 获取产品
func (mb *MacBookBuilder) GetProduct() ComputerProduct {
	fmt.Printf("computer:%+v", mb.computer)
	return mb.computer
}
