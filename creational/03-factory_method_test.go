package creational

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	var fm = &FactoryMethodMacBookProduct{}
	p := FactoryMethodMakeProduct(fm, "top", "intel i7", "mac")
	t.Logf("TestFactoryMethod computer: type = %v,cpu = %v,os = %v", p.Type, p.CPU, p.OS)
}
