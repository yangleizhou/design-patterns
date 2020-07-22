package creational

import "testing"

func TestSimpleFactory(t *testing.T) {
	sf := &SimpleFactory{}
	sf.CreatProduct("mac")
}
