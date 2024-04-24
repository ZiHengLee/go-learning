package array

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		first := i
		third := len(nums) - 1
		for second := i + 1; second < len(nums); second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third]+nums[first] > 0 {
				third--
			}
			if nums[first]+nums[second]+nums[third] == 0 {
				// 如果指针重合，随着 b 后续的增加
				// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
				if second == third {
					break
				}
				temp := []int{nums[first], nums[second], nums[third]}
				result = append(result, temp)
			}

		}
	}
	return result
}
