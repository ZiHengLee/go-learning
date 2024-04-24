package array

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	print(twoSum(nums, target))
}
