package structural

import "testing"

func TestDecorator(t *testing.T) {
	var man Man = &SingleMan{Name: "tom"}
	man = NewCarDecoratorImpl(man, "凯迪拉克 SUV")
	man = NewSchoolDecoratorImpl(man, "清华大学")
	man.GetHumanDesc()
}
