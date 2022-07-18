package mysort

import (
	"fmt"
	"testing"
)

func TestSortHalf(t *testing.T) {
	a := []int{5,1,1,2,0,0}

	quickSort(a,0,5)
	fmt.Println(a)
}
