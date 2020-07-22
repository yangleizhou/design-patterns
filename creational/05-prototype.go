package creational

//用一个已经创建的实例作为原型，通过复制该原型对象来创建一个和原型相同或相似的新对象，同时又能保证性能

//Cloneable 抽象原型类：规定了具体原型对象必须实现的接口
type Cloneable interface {
	Clone() Cloneable
}

//Type1 具体原型类,实现抽象原型类的 clone() 方法
type Type1 struct {
	name string
}

// Clone 复制新对象
func (t *Type1) Clone() Cloneable {
	ct := *t
	return &ct
}

//Type2 具体原型类,实现抽象原型类的 clone() 方
type Type2 struct {
	name string
}

// Clone 复制新对象
func (t *Type2) Clone() Cloneable {
	ct := *t
	return &ct
}

//原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象
//原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置

// PropotypeManager 原型管理器
type PropotypeManager struct {
	prototypes map[string]Cloneable
}

// Get 获取Cloneable
func (pm *PropotypeManager) Get(name string) Cloneable {
	return pm.prototypes[name]
}

// Set 根据name设置Cloneable
func (pm *PropotypeManager) Set(name string, ca Cloneable) {
	pm.prototypes[name] = ca
}

var pm PropotypeManager

func init() {
	pm.prototypes = make(map[string]Cloneable)
	pm.prototypes["t1"] = &Type1{name: "t1"}
}
