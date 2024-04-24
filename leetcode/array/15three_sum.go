package array

import (
	"sort"
)

// 1.前后指针一起移动
// 2.关键是重复元素跳过边界问题

func threeSum(nums []int) [][]int {
	lenNums := len(nums)
	sort.Ints(nums)
	ret := make([][]int, 0)
	// 这里的小于lenNums-2 是为了给left，right都留一个位置
	for i := 0; i < lenNums-2; i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		tempArr := make([]int, 0)
		left, right := i+1, lenNums-1
		tempArr = append(tempArr, nums[i])
		for left < right {
			//边界条件也很关键
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			} else if right+1 < lenNums && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[i]+nums[left]+nums[right] == 0 {
				helpArray := tempArr
				helpArray = append(helpArray, nums[left], nums[right])
				ret = append(ret, helpArray)
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ret
}
