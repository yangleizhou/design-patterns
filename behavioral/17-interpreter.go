package behaivoral

//Expression  定义解释器的接口，约定解释器的解释操作，主要包含解释方法 interpret()
type Expression interface {
	Interpreter()
}

//Context
type Context struct {
	val string
}
