package creational

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	p1 := pm.Get("t1")
	p2 := p1.Clone()
	fmt.Printf("p1 = %p,p2 = %p,p1.name = %v,p2.name = %v", p1, p2, p1, p2)
}
