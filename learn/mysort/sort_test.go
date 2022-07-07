package mysort

import (
	"fmt"
	"testing"
)

func TestSortHalf(t *testing.T) {
	a := []int{3,2,1,5,6,4}

	fmt.Println(heapSort(a,2))
}
