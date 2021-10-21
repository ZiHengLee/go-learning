package others

import (
	"fmt"
	"testing"
)
//nums1 := []int{1,2}
//nums2 := []int{1,2,3,4}
//func TestOthers(t *testing.T) {
//	fmt.Println(isValid("()[]{}"))
//	//fmt.Println(isValid("(()()())())"))
//}
func TestThreeNum(t *testing.T) {
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))
}
func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1,2}
	nums2 := []int{3,4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

