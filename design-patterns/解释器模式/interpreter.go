package 解释器模式

type Expression interface {
	Interpreter(a, b int) int
}

type AddExpression struct {
}

func (t AddExpression) Interpreter(a, b int) int {
	return a + b
}

type MulExpression struct {
}

func (t MulExpression) Interpreter(a, b int) int {
	return a * b
}

type Pythagoras struct {
	expr1 Expression
	expr2 Expression
}

// 用加和减够成一个勾股定理
func (t Pythagoras) interperter(a, b int) int {
	return t.expr1.Interpreter(t.expr2.Interpreter(a, a), t.expr2.Interpreter(b, b))
}
