package behavioral

//在不破坏封装性的前提下,捕获一个对象的内部状态，并在对象之外保存这个状态,以便以后当需要时能将该对象恢复到原先保存的状态
//通过备忘录类来专门存储对象状态， 当需要保存状态时， 管理者从发起人拿到状态，保存到备忘录中(即将一个对象的状态保存到另一个对象中)

// Originator 发起人
type Originator struct {
	Name string
}

// GetName 获取名字
func (o *Originator) GetName() string {
	return o.Name
}

// SetName 设置名字
func (o *Originator) SetName(name string) {
	o.Name = name
}

//NewMemento 创建备忘录
func (o *Originator) NewMemento() *Memento {
	return &Memento{Name: o.Name}
}

//RestoreMemento 恢复备忘录
func (o *Originator) RestoreMemento(m *Memento) {
	o.SetName(m.GetName())
}

// Memento 备忘录
type Memento struct {
	Name string
}

// GetName 获取名字
func (m *Memento) GetName() string {
	return m.Name
}

// SetName 设置名字
func (m *Memento) SetName(name string) {
	m.Name = name
}

// Caretaker 管理者
type Caretaker struct {
	list []*Memento
	top  int
}

// Pop 获取
func (ck *Caretaker) Pop() *Memento {
	if ck.top <= 0 {
		return nil
	}
	ck.top--
	return ck.list[ck.top]
}

// Push 保存
func (ck *Caretaker) Push(m *Memento) {
	ck.list = append(ck.list, m)
	ck.top++
}
