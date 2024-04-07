package 访问者模式

import "fmt"

type Visitor interface {
	BuyApple(a Apple)
	BuyBanana(b Banana)
}

type NormalVisitor struct {
}

func (t NormalVisitor) BuyApple(a Apple) {
	fmt.Println("苹果原价:", a.GetPrice(), "零售价格:", a.GetPrice())
}

func (t NormalVisitor) BuyBanana(b Banana) {
	fmt.Println("香蕉原价:", b.GetPrice(), "零售价格:", b.GetPrice())
}
