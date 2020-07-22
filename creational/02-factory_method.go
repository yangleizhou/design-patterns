package creational

import "fmt"

// 使一个类的实例化(具体产品创建)延迟到其子类(具体工厂)中完成, 定义一个用于创建对象的接口, 让子类决定将哪一个类实例化

// FactoryMethod 抽象工厂,定义具体工厂的公共接口
type FactoryMethod interface {
	Create() FactoryMethodProduct
}

// FactoryMethodProduct 抽象产品,定义具体产品的公共接口
type FactoryMethodProduct interface {
	SetCPU(cpu string)
	SetType(typ string)
	SetOS(os string)
	GetProduct() *ComputerProduct
}

// FactoryMethodMacBookProduct 具体产品类
// golang 没有继承,使用匿名组合
type FactoryMethodMacBookProduct struct {
	*ComputerProduct
}

// Create 具体工厂,定义创建对应具体产品实例的方法
func (mbp *FactoryMethodMacBookProduct) Create() FactoryMethodProduct {
	return &FactoryMethodMacBookProduct{
		ComputerProduct: &ComputerProduct{}}
}

// GetProduct 获取产品
func (mbp *FactoryMethodMacBookProduct) GetProduct() *ComputerProduct {
	fmt.Printf("mac 电脑配置:type = %v,cpu = %v,os = %v", mbp.Type, mbp.CPU, mbp.OS)
	return mbp.ComputerProduct
}

// FactoryMethodMakeProduct 工厂内部流程
func FactoryMethodMakeProduct(fm FactoryMethod, typ, cpu, os string) *ComputerProduct {
	obj := fm.Create()
	obj.SetCPU(cpu)
	obj.SetOS(os)
	obj.SetType(typ)
	return obj.GetProduct()
}
