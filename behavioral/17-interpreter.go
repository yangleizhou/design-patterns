package behavioral

import (
	"fmt"
	"strconv"
	"strings"
)

//Expression  定义解释器的接口，约定解释器的解释操作，主要包含解释方法 interpret()
type Expression interface {
	Interpreter() int
	Print()
}

//Context 上下文
type Context struct {
	val int
}

// Interpreter 获取值
func (c *Context) Interpreter() int {
	return c.val
}

func (c *Context) Print() {
	fmt.Printf("%d", c.val)
}

//AddExpression 表达式
type AddExpression struct {
	left, right Expression
	deepth      string //打印日志用
}

// Interpreter 实现
func (add *AddExpression) Interpreter() int {
	return add.left.Interpreter() + add.right.Interpreter()
}

// Print
func (add *AddExpression) Print() {
	add.left.Print()
	fmt.Print("+")
	add.right.Print()
}

//SubExpression 表达式
type SubExpression struct {
	left, right Expression
}

// Interpreter 实现
func (sub *SubExpression) Interpreter() int {
	return sub.left.Interpreter() - sub.right.Interpreter()
}

// Print da
func (sub *SubExpression) Print() {
	sub.left.Print()
	fmt.Print("-")
	sub.right.Print()
}

//Parser 要任务是将需要分析的句子或表达式转换成使用解释器对象描述的抽象语法树
type Parser struct {
	exp []string
	pre Expression
}

// Parse 解析句子成抽象语法树
func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")
	var len = len(p.exp)
	var deepth string
	for i := 0; i < len; i++ {
		switch p.exp[i] {
		case "+":
			i++
			p.pre = p.newAdd(i)
		case "-":
			i++
			p.pre = p.newSub(i)
		default:
			p.pre = p.Val(i)
		}
		deepth += " "
	}
}

// Print 打印语法树
func (p *Parser) Print() {
	p.pre.Print()
}

//GetResult 获取结构
func (p *Parser) GetResult() {
	p.Print()
	fmt.Printf(" = %d\n", p.pre.Interpreter())
}

// newAdd 获取Add Expression
func (p *Parser) newAdd(i int) Expression {
	val, _ := strconv.Atoi(p.exp[i])
	return &AddExpression{left: p.pre, right: &Context{val: val}}
}

// newSub 获取Sub Expression
func (p *Parser) newSub(i int) Expression {
	val, _ := strconv.Atoi(p.exp[i])
	return &SubExpression{left: p.pre, right: &Context{val: val}}
}

// Val 获取Context Expression
func (p *Parser) Val(i int) Expression {
	val, _ := strconv.Atoi(p.exp[i])
	return &Context{val: val}
}
