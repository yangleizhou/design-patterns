package behavioral

import "testing"

func TestInterpreter(t *testing.T) {
	parse := &Parser{}
	parse.Parse("1 + 2 + 3 - 8")
	parse.GetResult()
}
