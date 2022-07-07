package bilinknode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestLRU(t *testing.T) {
	a := "皮友21312312a"
	if strings.HasPrefix(a, "皮友") {
		_, numErr := strconv.Atoi(strings.ReplaceAll(a, "皮友", ""))
		if numErr == nil {
			print("haha")
		}

	}
}
func ListPrint(node *LinkNode) {
	for i := node; i != nil; {
		fmt.Print(i.val)
		fmt.Print("-")
		i = i.next
	}
	fmt.Println()
}
