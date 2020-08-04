package behavioral

import "fmt"

//定义了一系列算法，并将每个算法封装起来，使它们可以相互替换，且算法的变化不会影响使用算法的客户。
//策略模式属于对象行为模式，它通过对算法进行封装，把使用算法的责任和算法的实现分割开来，并委派给不同的对象对这些算法进行管理,符合开闭原则。

var _ Strategy = (*ConcreteStrategyA)(nil)
var _ Strategy = (*ConcreteStrategyB)(nil)

//Strategy 抽象策略类，定义了一个公共接口，各种不同的算法以不同的方式实现这个接口，
//环境角色使用这个接口调用不同的算法，一般使用接口或抽象类实现
type Strategy interface {
	Algorithm()
}

//ContextStrategy 环境类,持有一个策略类的引用，最终给客户端调用
type ContextStrategy struct {
	s Strategy
}

// SetStrategy 设置策略
func (c *ContextStrategy) SetStrategy(s Strategy) {
	c.s = s
}

//Algorithm 调用策略算法
func (c *ContextStrategy) Algorithm() {
	c.s.Algorithm()
}

// ConcreteStrategyA 具体策略类，实现了抽象策略定义的接口，提供具体的算法实现
type ConcreteStrategyA struct {
}

//Algorithm 具体算法实现
func (s ConcreteStrategyA) Algorithm() {
	fmt.Println("use a algorithm......")
}

// ConcreteStrategyB 具体策略类，实现了抽象策略定义的接口，提供具体的算法实现
type ConcreteStrategyB struct {
}

//Algorithm 具体算法实现
func (s ConcreteStrategyB) Algorithm() {
	fmt.Println("use b algorithm......")
}
