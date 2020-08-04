package structural

import (
	"fmt"
)

//将对象组合成树状层次结构，使用户对单个对象和组合对象具有一致的访问性

var _ Component = (*Leaf)(nil)
var _ Component = (*Composite)(nil)

//node 节点类型
const (
	CompositeNode = iota //树枝
	LeafNode             //树叶

)

//Component 抽象构件角色,实现所有类共有接口的默认行为
type Component interface {
	SetParent(Component)
	GetParent() Component
	AddChild(Component)
	SetPath(path string) Component
	GetPath() string
	SetNodeType(nodeType int) Component
	GetNodeType() int
	Print(pre string)
}

//NewComponent 获得构件
func NewComponent(nodeType int, path string) Component {
	var node Component
	switch nodeType {
	case LeafNode:
		node = &Leaf{}
	case CompositeNode:
		node = &Composite{
			childs: make([]Component, 0),
		}
	default:
		panic("never reach here")
	}
	node.SetNodeType(nodeType).SetPath(path)
	return node
}

//component 节点对象
type component struct {
	parent   Component
	Path     string
	NodeType int
}

//SetParent 设置parent
func (c *component) SetParent(parent Component) {
	c.parent = parent
}

//GetParent 所有节点类型通用
func (c *component) GetParent() Component {
	return c.parent
}

//SetNodeType 设置节点类型
func (c *component) SetNodeType(nodeType int) Component {
	c.NodeType = nodeType
	return c
}

//GetNodeType 获取节点类型
func (c *component) GetNodeType() int {
	return c.NodeType
}

//SetPath 设置path
func (c *component) SetPath(path string) Component {
	c.Path = path
	return c
}

//GetPath 获取节点Path
func (c *component) GetPath() string {
	return c.Path
}

//AddChild 增加子节点
func (c *component) AddChild(Component) {

}

//Print 打印信息
func (c *component) Print(pre string) {

}

//Leaf 树叶构件角色
type Leaf struct {
	component
}

//Print 叶子节点打印
func (node *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, node.GetPath())
}

//Composite 树枝构件角色
type Composite struct {
	component
	childs []Component
}

//AddChild 增加子节点
func (node *Composite) AddChild(child Component) {
	if node.CheckRing(child) {
		fmt.Println("add child lead to ring ")
		return
	}
	child.SetParent(node)
	node.childs = append(node.childs, child)
}

//Print 树枝打印
func (node *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, node.GetPath())
	pre += " "
	for _, comp := range node.childs {
		comp.Print(pre)
	}
}

//CheckRing 添加前判断是否成环
func (node *Composite) CheckRing(child Component) bool {
	if node.GetNodeType() == child.GetNodeType() &&
		node.GetPath() == child.GetPath() {
		return true
	}
	p := node.GetParent()
	if p == nil {
		return false
	}
	p.(*Composite).CheckRing(child)
	return false
}
