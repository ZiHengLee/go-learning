package others
//https://leetcode.cn/problems/find-the-duplicate-number/
//寻找重复数，有点像老筋急转弯
func findDuplicate(nums []int) int {
	slow,fast := 0,0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow{
			fast = 0
			for nums[fast]!=nums[slow]{
				fast = nums[fast]
				slow = nums[slow]
			}
			return nums[fast]
		}
	}
}
