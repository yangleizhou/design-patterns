package behavioral

import "testing"

func TestCommand(t *testing.T) {
	invoker := &Invoker{}
	rev := &receiver{}
	c1 := &command{Receiver: rev, Button: buttonOpen}
	c2 := &command{Receiver: rev, Button: buttonClose}
	invoker.AddCommand(c1, c2)
	invoker.ExecuteCommand()
}
