package structural

import "testing"

func TestFlyweight(t *testing.T) {
	// 享元工厂对象
	var f FlyweightFactory
	f.GetFlyweight("one").Get()
	f.GetFlyweight("one").Get()
	f.GetFlyweight("two").Get()
}
