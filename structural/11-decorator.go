package structural

//Component 抽象构件
//定义一个对象接口或抽象类，以规范需要装饰的主体类，比如：人
type Component interface {
	GetHumanDesc()
}

//NormalMan  男生
type NormalMan struct {
}
