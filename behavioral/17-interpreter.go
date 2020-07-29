package behavioral

import "strings"

//Expression  定义解释器的接口，约定解释器的解释操作，主要包含解释方法 interpret()
type Expression interface {
	Interpreter() int
}

//Context 上下文
type Context struct {
	val int
}

// Interpreter 获取值
func (c *Context) Interpreter() int {
	return c.val
}

//AddExpression 表达式
type AddExpression struct {
	left, right Expression
}

// Interpreter 实现
func (add *AddExpression) Interpreter() int {
	return add.left.Interpreter() + add.right.Interpreter()
}

//SubExpression 表达式
type SubExpression struct {
	left, right Expression
}

// Interpreter 实现
func (sub *SubExpression) Interpreter() int {
	return sub.left.Interpreter() - sub.right.Interpreter()
}

//MulExpression 表达式
type MulExpression struct {
	left, right Expression
}

//Interpreter 实现
func (mul *MulExpression) Interpreter() int {
	return mul.left.Interpreter() * mul.right.Interpreter()
}

//Parser 要任务是将需要分析的句子或表达式转换成使用解释器对象描述的抽象语法树
type Parser struct {
	exp []string
	pre Expression
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")
	var len = len(p.exp)

	for i := 0; i < len; i++ {
		switch p.exp[i] {

		}
	}
}
