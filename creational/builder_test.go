package creational

import (
	"testing"
)

func TestBuilderFactory(t *testing.T) {
	d1 := &Director{}
	mac := &MacBookBuilder{}
	d1.MakeProduct(mac)
	p := mac.GetProduct()
	t.Logf("TestBuilderFactory computer: type = %v,cpu = %v,os = %v", p.Type, p.CPU, p.OS)
}
