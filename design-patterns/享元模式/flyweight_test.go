package 享元模式

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	fmt.Println(Sentence([]int{1, 2, 3, 4}))
	fmt.Println(Sentence([]int{3, 4, 2, 1}))
}
