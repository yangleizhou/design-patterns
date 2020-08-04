package behavioral

import "testing"

func TestIterator(t *testing.T) {
	var ag Aggregate = &Container{}
	ag.Add(&Teacher{})
	ag.Add(&Analysis{})
	iter := ag.Iterator()
	for iter.First(); iter.HasNext(); iter.Next() {
		iter.Value()
	}
}
