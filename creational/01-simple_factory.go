package creational

import "fmt"

var _SimpleIComputerProduct = (*SimpleMacBookProduct)(nil)

// SimpleFactory 工厂角色
type SimpleFactory struct {
}

// CreatProduct 创建产品
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
	fmt.Printf("computer:%+v", smp.computer)
	return smp.computer
}

// NewSimpleComputerProduct 创建对象
func NewSimpleComputerProduct() *SimpleMacBookProduct {
	return &SimpleMacBookProduct{}
}
