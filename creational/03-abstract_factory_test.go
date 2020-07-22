package creational

import "testing"

func TestAbstractFactory1(t *testing.T) {
	var af MacBookFactory
	af.FactoryComputer().GetProduct()
	af.FactoryPhone().GetProduct()
}

func TestAbstractFactory2(t *testing.T) {
	var af IOSPhoneFactory
	af.FactoryComputer().GetProduct()
	af.FactoryPhone().GetProduct()
}
