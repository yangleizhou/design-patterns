package structural

import "testing"

func TestFacade(t *testing.T) {
	s := NewBuyService()
	s.Service()
}
