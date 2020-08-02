package behavioral

import "testing"

func TestState(t *testing.T) {
	c := &Context{}
	s := &state{ctx: c}

	c.SetState(newLowState(s))
	c.Handle(50)
	c.Handle(50)
	c.Handle(-20)
}
