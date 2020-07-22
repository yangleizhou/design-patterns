package creational

//将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示

var _ Builder = (*MacBookBuilder)(nil)

// Builder 构建者的抽象基类
type Builder interface {
	SetType() Builder
	SetCPU() Builder
	SetOS() Builder
	GetProduct() ComputerProduct
}

// Director 指挥者,负责组装方法
type Director struct {
	builder Builder
}

// MakeProduct 组装
func (d *Director) MakeProduct(builder Builder) {
	d.builder = builder
	d.builder.SetType().SetCPU().SetOS()
}

// MacBookBuilder (具体建造者)苹果笔记本
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
	return mb.computer
}
