package 访问者模式

type Fruit interface {
	GetPrice() int
}

// Apple 苹果类
type Apple struct {
	price int
}

func (t Apple) GetPrice() int {
	return t.price
}

// Banana 香蕉类
type Banana struct {
	price int
}

func (t Banana) GetPrice() int {
	return t.price
}
