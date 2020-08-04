package behavioral

import "fmt"

//对有状态的对象，把复杂的“判断逻辑”提取到不同的状态对象中，允许状态对象在其内部状态发生改变时改变其行为。

var _ State = (*state)(nil)
var _ State = (*LowState)(nil)
var _ State = (*MidState)(nil)
var _ State = (*HighState)(nil)

//Context 上下文，它定义了客户感兴趣的接口，维护一个当前状态，并将与状态相关的操作委托给当前状态对象来处理
type Context struct {
	state State
}

//Handle 处理
func (c *Context) Handle(score int) {
	c.state.Handle(score)
}

// SetState 设置状态
func (c *Context) SetState(state State) {
	c.state = state
}

// GetState 获取状态
func (c *Context) GetState() State {
	return c.state
}

// State 抽象状态
type State interface {
	Handle(score int)
	ChangeState()
}

// state 具体状态
type state struct {
	ctx       *Context
	score     int
	levelName string
}

// Handle 状态处理
func (s *state) Handle(score int) {
	if score >= 0 {
		fmt.Printf("score %d + %d 分", s.score, score)
	} else {
		fmt.Printf("score %d - %d 分", s.score, score)
	}
	s.score += score
	if s.score > 100 {
		s.score = 100
	}
	if s.score < 0 {
		s.score = 0
	}
	s.ctx.state.ChangeState()
	fmt.Printf(" %s 哦\n ", s.levelName)
}

// ChangeState 状态改变逻辑,复杂的“判断逻辑”提取到不同的状态对象中
func (s *state) ChangeState() {
}

// LowState 低分
// golang 没有继承使用匿名组合
type LowState struct {
	*state
}

// ChangeState 状态变更判断
func (ls *LowState) ChangeState() {
	var s State
	if ls.score >= 90 {
		s = newHighState(ls.state)
	} else if ls.score >= 60 {
		s = newMidState(ls.state)
	} else {
		return
	}
	ls.ctx.SetState(s)
}

// newLowState 获取低分对象
func newLowState(s *state) State {
	s.levelName = "不及格"
	return &LowState{state: s}
}

// MidState 中等
type MidState struct {
	*state
}

// ChangeState 状态变更判断
func (ms *MidState) ChangeState() {
	var s State
	if ms.score < 60 {
		s = newLowState(ms.state)
	} else if ms.score >= 90 {
		s = newHighState(ms.state)
	} else {
		return
	}
	ms.ctx.SetState(s)
}

// newMidState 获取中分对象
func newMidState(s *state) State {
	s.levelName = "中等"
	return &MidState{state: s}
}

// HighState 高分
type HighState struct {
	*state
}

// ChangeState 状态变更判断
func (hs *HighState) ChangeState() {
	var s State
	if hs.score < 60 {
		s = newLowState(hs.state)
	} else if hs.score < 90 {
		s = newMidState(hs.state)
	} else {
		return
	}
	hs.ctx.SetState(s)
}

// newHighState 获取高分对象
func newHighState(s *state) State {
	s.levelName = "优秀"
	return &HighState{state: s}
}
