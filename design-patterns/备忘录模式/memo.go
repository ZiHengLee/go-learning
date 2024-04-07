package 备忘录模式

import "strconv"

type Memo interface {
	GetState() string
}

type CalculatorI interface {
	Add(int)
	Sub(int)
	GetResult() int
	Save() Memo
	Restore(Memo)
}

type calculatorMemo struct {
	state string
}

func (c *calculatorMemo) GetState() string {
	return c.state
}

type Calculator struct {
	result int
}

func (c *Calculator) Add(x int) {
	c.result += x
}

func (c *Calculator) Sub(x int) {
	c.result -= x
}

func (c *Calculator) GetResult() int {
	return c.result
}

func (c *Calculator) Save() Memo {
	return &calculatorMemo{state: strconv.Itoa(c.result)}
}

func (c *Calculator) Restore(m Memo) {
	if m != nil {
		if cm, ok := m.(*calculatorMemo); ok {
			if result, err := strconv.Atoi(cm.state); err == nil {
				c.result = result
			}
		}
	}
}
