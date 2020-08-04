package behavioral

import "testing"

func TestMediator(t *testing.T) {
	var c Colleague = &colleague{}
	var m = &Market{Colleague: c}
	var tech = &Technical{Colleague: c}
	mediator := NewMediator()
	mediator.Registered(colleagueTypeMarket, m)
	mediator.Registered(colleagueTypeTech, tech)
	mediator.Operation(colleagueTypeMarket, "市场部需要招人->人事")
	mediator.Operation(colleagueTypeTech, "人事->技术部下午发年终")
}
