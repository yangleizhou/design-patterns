package creational

//提供一个创建一系列相关或相互依赖对象的接口, 而无需指定它们具体的类

var _ AbstractFactory = (*MacBookFactory)(nil)
var _ AbstractFactory = (*IOSPhoneFactory)(nil)

// AbstractComputer 创建抽象产品类
type AbstractComputer interface {
	GetProduct()
}

// AbstractPhone 创建抽象产品类
type AbstractPhone interface {
	GetProduct()
}

// AbstractFactory 抽象工厂,定义具体工厂的公共接口
type AbstractFactory interface {
	FactoryComputer() AbstractComputer
	FactoryPhone() AbstractPhone
}

// MacBookFactory 具体工厂
type MacBookFactory struct {
}

// FactoryComputer 定义创建对应具体产品实例的方法
func (*MacBookFactory) FactoryComputer() AbstractComputer {
	return &ComputerProduct{}
}

// FactoryPhone 定义创建对应具体产品实例的方法
func (*MacBookFactory) FactoryPhone() AbstractPhone {
	return &PhoneProduct{}
}

// IOSPhoneFactory 具体工厂
type IOSPhoneFactory struct {
}

// FactoryComputer 定义创建对应具体产品实例的方法
func (*IOSPhoneFactory) FactoryComputer() AbstractComputer {
	return &ComputerProduct{}
}

// FactoryPhone 定义创建对应具体产品实例的方法
func (*IOSPhoneFactory) FactoryPhone() AbstractPhone {
	return &PhoneProduct{}
}
