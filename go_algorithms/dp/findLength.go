package dp

//最长重复子数组
//思路
//dp[i][j]代表以A[i-1]与B[j-1]结尾的公共字串的长度
func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0|| len(nums2) == 0{
		return 0
	}
	dp := make([][]int,len(nums1)+1)
	maxLen := 0
	dp[0] = make([]int,len(nums2)+1)
	for i:=1;i<len(nums1)+1;i++{
		dp[i] = make([]int,len(nums2)+1)
		for j:=1;j<len(nums2)+1;j++{
			if nums1[i-1] == nums2[j-1]{
				dp[i][j] = dp[i-1][j-1]+1
			}
			if dp[i][j] > maxLen{
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen
}
