package behavioral

import "testing"

func TestStrategy(t *testing.T) {
	var s Strategy = &ConcreteStrategyA{}
	c := &ContextStrategy{}
	c.SetStrategy(s)
	c.Algorithm()

	s = &ConcreteStrategyB{}
	c.SetStrategy(s)
	c.Algorithm()
}
