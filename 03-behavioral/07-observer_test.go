package behavioral

import "testing"

func TestObserver(t *testing.T) {
	investor1 := NewInvestorObserver("特普朗")
	investor2 := NewInvestorObserver("乔布斯")
	investor3 := NewInvestorObserver("马云")

	var ob Subject = &subject{}
	ob.Register(investor1).Register(investor2).Register(investor3)
	ob.Notify("股票下跌啦")

	ob.Remove(investor1)

	ob.Notify("XXX宣布破产了")
}
