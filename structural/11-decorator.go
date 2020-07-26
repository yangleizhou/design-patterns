package structural

import "fmt"

//在不改变现有对象结构情况下(在不影响其他对象的情况下),动态地给一个对象增加一些职责，即增加其额外的功能

var _ Man = (*SingleMan)(nil)
var _ AttachePropertiesDecorator = (*CarDecoratorImpl)(nil)
var _ AttachePropertiesDecorator = (*SchoolDecoratorImpl)(nil)

//Man 抽象构件 - Component
//定义一个对象接口或抽象类，以规范需要装饰的主体类，比如：人
type Man interface {
	GetHumanDesc()
}

//SingleMan  单身Boy - ConcreteComponent
// 具体构件 实际被动态地添加职责的对象
type SingleMan struct {
	Name string
}

//GetHumanDesc 实现GetHumanDesc接口实现
func (sm *SingleMan) GetHumanDesc() {
	fmt.Println(sm.Name)
}

//AttachePropertiesDecorator 附加属性接口 -  Decorator
type AttachePropertiesDecorator interface {
	AddAttachProperties(desc string)
}

// CarDecoratorImpl 小车装饰者 - ConcreteDecorator
type CarDecoratorImpl struct {
	Desc string
	m    Man
}

// AddAttachProperties 小车属性
func (car *CarDecoratorImpl) AddAttachProperties(desc string) {
	car.Desc = desc
}

// GetHumanDesc 动态地添加职责的对象
func (car *CarDecoratorImpl) GetHumanDesc() {
	car.m.GetHumanDesc()
	fmt.Println("car:", car.Desc)
}

//NewCarDecoratorImpl 车装饰者对象
func NewCarDecoratorImpl(m Man, desc string) Man {
	var car = &CarDecoratorImpl{m: m}
	car.AddAttachProperties(desc)
	return car
}

// SchoolDecoratorImpl 学校装饰者 - ConcreteDecorator
type SchoolDecoratorImpl struct {
	Desc string
	m    Man
}

//AddAttachProperties 学校属相
func (school *SchoolDecoratorImpl) AddAttachProperties(desc string) {
	school.Desc = desc
}

// GetHumanDesc 动态地添加职责的对象
func (school *SchoolDecoratorImpl) GetHumanDesc() {
	school.m.GetHumanDesc()
	fmt.Println("school:", school.Desc)
}

//NewSchoolDecoratorImpl 学校装饰者对象
func NewSchoolDecoratorImpl(m Man, desc string) Man {
	var shool = &SchoolDecoratorImpl{m: m}
	shool.AddAttachProperties(desc)
	return shool
}
