package others

//和为k的连续子数组的个数，
//https://leetcode.cn/problems/subarray-sum-equals-k/submissions/

//有点像取巧题
//使用前缀数组的思想，保存每个前缀和可能出现的个数，当出现一个当下前缀和a-k=之前的前缀和b时，其实b的个数不就是这种连续数组出现的个数吗
//但要注意初始条件map[0]=1
func subarraySum(nums []int, k int) int {
	helpMap := map[int]int{0: 1}
	subSum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		subSum += nums[i]
		if helpMap[subSum-k] > 0 {
			res += helpMap[subSum-k]
		}
		helpMap[subSum]+=1
	}
	return res
}
