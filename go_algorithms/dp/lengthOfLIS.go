package dp

//获取严格递增子序列长度

//用最后一位结尾来标注dp
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) < 2{
		return len(nums)
	}
	dp[0] = 1
	curMax := 0
	for i:=1;i<len(nums);i++{
		dp[i] = 1
		for j:=0;j<i;j++{
			if nums[i] > nums[i-1]{
				if dp[j]+1 > dp[i]{
					dp[i] = dp[j]+1
				}
			}
		}
		if dp[i] > curMax{
			curMax = dp[i]
		}
	}
	return curMax
}
