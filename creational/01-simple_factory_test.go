package creational

import (
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	sf := &SimpleFactory{}
	p := sf.CreatProduct("mac").GetProduct()
	t.Logf("TestSimpleFactory computer: type = %v,cpu = %v,os = %v", p.Type, p.CPU, p.OS)
}
