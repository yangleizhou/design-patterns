package creational

//又称为静态工厂方法(Static Factory Method)模式，它属于类创建型模式。在简单工厂模式中，可以根据参数的不同返回不同类的实例。
//简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。

var _SimpleIComputerProduct = (*SimpleMacBookProduct)(nil)

// SimpleFactory 工厂角色
type SimpleFactory struct {
}

// CreatProduct 根据参数,创建产品
func (sf *SimpleFactory) CreatProduct(os string) SimpleIComputerProduct {
	switch os {
	case "mac":
		cp := NewSimpleComputerProduct()
		MakeProduct(&cp.computer,
			SetType("top"),
			SetCPU("intel i7"),
			SetOS(os))
		return cp
	}
	return nil
}

// SimpleIComputerProduct 抽象产品角色
type SimpleIComputerProduct interface {
	GetProduct() ComputerProduct
}

// SimpleMacBookProduct 具体产品角色1-苹果电脑
type SimpleMacBookProduct struct {
	computer ComputerProduct
}

// GetProduct 获取苹果电脑
func (smp *SimpleMacBookProduct) GetProduct() ComputerProduct {
	return smp.computer
}

// NewSimpleComputerProduct 创建对象
func NewSimpleComputerProduct() *SimpleMacBookProduct {
	return &SimpleMacBookProduct{}
}
