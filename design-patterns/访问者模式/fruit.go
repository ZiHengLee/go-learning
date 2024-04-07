package 访问者模式

type NewFruit interface {
	Accept(v NewVisitor)
	GetPrice() int
}

type NewApple struct {
	price int
}

func (t NewApple) Accept(v NewVisitor) {
	v.BuyApple(t)
}
func (t NewApple) GetPrice() int {
	return t.price
}

type NewBanana struct {
	price int
}

func (t NewBanana) Accept(v NewVisitor) {
	v.BuyBanana(t)
}
func (t NewBanana) GetPrice() int {
	return t.price
}
