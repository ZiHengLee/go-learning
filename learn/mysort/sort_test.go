package mysort

import (
	"fmt"
	"testing"
)

func TestSortHalf(t *testing.T) {
	nums := []int{1,3,1,4,2,85,9}
	quickSort(nums,0,6)
	fmt.Println(nums)
}
