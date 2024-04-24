package array

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp1 := target - nums[i]
		if _, ok := tmp[tmp1]; ok {
			tmp2 := tmp[tmp1]
			return []int{tmp2, i}
		}
		tmp[nums[i]] = i
	}
	return nil
}
