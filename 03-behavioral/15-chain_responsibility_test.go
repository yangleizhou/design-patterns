package behavioral

import "testing"

func TestChainResponsibility(t *testing.T) {
	bossChain := NewRequestChain(chainBoss)
	depChain := NewRequestChain(chainDep).SetNexChain(bossChain)
	chain := NewRequestChain(chainProj).SetNexChain(depChain)

	//funding 500
	chain.Handle(500)
	chain.Handle(10000)
	chain.Handle(4000)
}
