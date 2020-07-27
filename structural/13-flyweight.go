package structural

import "fmt"

//运用共享技术来有效地支持大量细粒度对象的复用

var _ Flyweight = (*Circle)(nil)

//FlyweightFactory 享元工厂类
type FlyweightFactory struct {
	mgr map[string]Flyweight
}

//GetFlyweight 获取享元池对象
func (f *FlyweightFactory) GetFlyweight(k string) Flyweight {
	if f.mgr == nil {
		f.mgr = make(map[string]Flyweight)
	}
	return f.Get(k)
}

//Get 获取享元对象
func (f *FlyweightFactory) Get(k string) Flyweight {
	if obj, ok := f.mgr[k]; ok {
		return obj
	}
	return &Circle{color: k}

}

//Flyweight 抽象享元类
type Flyweight interface {
	Get()
}

//Circle 具体享元类
type Circle struct {
	color string
}

//Get 实现抽象享元类接口
func (c *Circle) Get() {
	fmt.Println("circle addr = ", &c)
}
