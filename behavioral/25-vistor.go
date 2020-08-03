package behavioral

import "fmt"

//将作用于某种数据结构中的各元素的操作分离出来封装成独立的类，
//使其在不改变数据结构的前提下可以添加作用于这些元素的新的操作，
//为数据结构中的每个元素提供多种访问方式,它将对数据的操作与数据结构进行分离

var _ Visitor = (*CustomerVisitorA)(nil)
var _ Visitor = (*CustomerVisitorB)(nil)
var _ Visitor = (*CashierVisitor)(nil)
var _ Element = (*ProductElement)(nil)

//Visitor 抽象访问者角色,定义一个访问具体元素的接口，为每个具体元素类对应一个访问操作 visit() ，
//该操作中的参数类型标识了被访问的具体元素
type Visitor interface {
	Visit(*ProductElement)
}

// CustomerVisitorA 消费者A访问者
type CustomerVisitorA struct {
	Name string
}

// Visit 实现 Visit
func (v *CustomerVisitorA) Visit(pe *ProductElement) {
	if pe.Prize >= 200 {
		fmt.Printf("商品:%s价格:%d超出顾客%s的预算\n", pe.Name, pe.Prize, v.Name)
		return
	}
	pe.BuyTag = true
	fmt.Printf("顾客 %s 选择了%s,因为性价比高\n", v.Name, pe.Name)
}

// CustomerVisitorB 消费者B访问者
type CustomerVisitorB struct {
	Name string
}

// Visit 实现 Visit
func (v *CustomerVisitorB) Visit(pe *ProductElement) {
	if pe.Prize < 100 {
		fmt.Printf("商品:%s价格:%d低于顾客%s的预算\n", pe.Name, pe.Prize, v.Name)
		return
	}
	pe.BuyTag = true
	fmt.Printf("顾客 %s 选择了%s,质量杠杠的!\n", v.Name, pe.Name)
}

// CashierVisitor 收费员访问者
type CashierVisitor struct {
	Name        string
	TotaolPrice int
}

// Visit 实现 Visit
func (v *CashierVisitor) Visit(pe *ProductElement) {
	if !pe.BuyTag {
		return
	}
	fmt.Printf("商品:%s --- 售价:%d\n", pe.Name, pe.Prize)
	v.TotaolPrice += pe.Prize
	pe.BuyTag = false
}

// Print 打印小票
func (v *CashierVisitor) Print() {
	fmt.Printf("您好,收费员%s为您服务,您购买商品共计%d,谢谢惠顾!\n\n\n", v.Name, v.TotaolPrice)
	v.TotaolPrice = 0
}

// Element 抽象元素角色,声明一个包含接受操作 accept() 的接口，被接受的访问者对象作为 accept() 方法的参数。
type Element interface {
	Accept(v Visitor)
}

// ProductElement 商品
type ProductElement struct {
	Name   string
	Prize  int
	BuyTag bool
}

// Accept Accept 接口实现
func (pe *ProductElement) Accept(v Visitor) {
	v.Visit(pe)
}

// newProductElement 获取商品
func newProductElement(name string, prize int) Element {
	return &ProductElement{Name: name, Prize: prize}
}

//ElementContainer 对象结构角色
type ElementContainer struct {
	list []Element
}

//Add 实现容器的元素的添加和移除
func (container *ElementContainer) Add(element Element) *ElementContainer {
	if container == nil || element == nil {
		return nil
	}
	container.list = append(container.list, element)
	return container
}

// Delete 移除Element
func (container *ElementContainer) Delete(element Element) {
	for i, val := range container.list {
		if val == element {
			container.list = append(container.list[:i], container.list[i+1:]...)
			break
		}
	}
}
