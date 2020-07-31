package behavioral

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	originator := &Originator{}
	caretaker := &Caretaker{}
	originator.SetName("aa")
	fmt.Println("name = ", originator.GetName())
	caretaker.Push(originator.NewMemento())

	originator.SetName("AA")
	fmt.Println("name = ", originator.GetName())

	memento := caretaker.Pop()
	originator.RestoreMemento(memento)
	fmt.Println("name = ", originator.GetName())

}
