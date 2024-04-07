package 访问者模式

import "fmt"

type NewVisitor interface {
	BuyApple(a NewApple)
	BuyBanana(b NewBanana)
}

type NewNormalVisitor struct {
}

func (t NewNormalVisitor) BuyApple(a NewApple) {
	fmt.Println("苹果原价:", a.GetPrice(), "零售价格:", a.GetPrice())
}

func (t NewNormalVisitor) BuyBanana(b NewBanana) {
	fmt.Println("香蕉原价:", b.GetPrice(), "零售价格:", b.GetPrice())
}
