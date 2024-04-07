package dp

//最大连续子序列和
//动态规划
//dp[i] = max(a[i],dp[i-1]+a[i])

func maxSubArray(nums []int) int {
	dp := make([]int,len(nums))
	maxCur := nums[0]
	for k,v := range nums{
		if k > 0{
			if v+dp[k-1] > v{
				dp[k] = v+dp[k-1]
			}else{
				dp[k] = v
			}
		}else{
			dp[k] = v
		}
		if dp[k] > maxCur{
			maxCur = dp[k]
		}
	}
	return maxCur
}