package behavioral

import (
	"fmt"
)

//通过将聚合对象的遍历行为分离出来，抽象成迭代器类来实现的，其目的是在不暴露聚合对象的内部结构的情况下，让外部代码透明地访问聚合的内部数据

var _ Aggregate = (*Container)(nil)
var _ Iterator = (*iterator)(nil)
var _ Visitor = (*Teacher)(nil)
var _ Visitor = (*Analysis)(nil)

//Aggregate 抽象聚合角色，定义存储、添加、删除聚合对象以及创建迭代器对象的接口
type Aggregate interface {
	Iterator() Iterator
	Add(v Visitor)
	Remove(index int)
}

//Container 具体聚合角色,实现抽象聚合类，返回一个具体迭代器的实例
type Container struct {
	list []Visitor
}

//Iterator 获取迭代器
func (c *Container) Iterator() Iterator {
	return newIterator(c)
}

//Add 添加
func (c *Container) Add(v Visitor) {
	c.list = append(c.list, v)
}

//Remove 移除
func (c *Container) Remove(index int) {
	if index < 0 || index >= len(c.list) {
		return
	}
	c.list = append(c.list[:index], c.list[index+1:]...)
}

//Iterator 抽象迭代器角色,定义访问和遍历聚合元素的接口
//通常包含 hasNext()、first()、next() 等方法
type Iterator interface {
	HasNext() bool
	First()
	Next()
	Value()
}

type iterator struct {
	pos int //记录当前访问位置
	*Container
}

// NewIterator 获取迭代器对象
func newIterator(container *Container) Iterator {
	iter := &iterator{}
	iter.Container = container
	return iter
}

//HasNext 是否有下一个
func (iter *iterator) HasNext() bool {
	if iter.pos < 0 || iter.pos >= len(iter.list) {
		return false
	}
	return true
}

// First 获取第一个
func (iter *iterator) First() {
	if len(iter.list) == 0 {
		iter.pos = -1
	}
}

// Next 获取下一个
func (iter *iterator) Next() {
	iter.pos++
}

// Value 获取迭代器的值
func (iter *iterator) Value() {
	iter.list[iter.pos].Visit()
}

// Visitor 对象接口
type Visitor interface {
	Visit()
}

//Teacher visitor对象
type Teacher struct{}

//Visit 实现接口
func (t *Teacher) Visit() {
	fmt.Println("this is teacher visitor")
}

//Analysis visitor对象
type Analysis struct{}

//Visit 实现接口
func (a *Analysis) Visit() {
	fmt.Println("this is analysis visitor")
}
