package behavioral

import "testing"

func TestVisitor(t *testing.T) {
	pe1 := newProductElement("九阳豆浆机", 400)
	pe2 := newProductElement("冷风机空调", 1000)
	pe3 := newProductElement("旺仔牛奶", 26)
	pe4 := newProductElement("boost shoe", 150)

	containter := &ElementContainer{}
	containter.Add(pe1).Add(pe2).Add(pe3).Add(pe4)

	cashier := &CashierVisitor{Name: "C"}

	v1 := &CustomerVisitorA{Name: "A"}
	v2 := &CustomerVisitorB{Name: "B"}

	for _, pe := range containter.list {
		pe.Accept(v1)
	}
	for _, pe := range containter.list {
		pe.Accept(cashier)
	}
	cashier.Print()

	for _, pe := range containter.list {
		pe.Accept(v2)
	}
	for _, pe := range containter.list {
		pe.Accept(cashier)
	}
	cashier.Print()

}
