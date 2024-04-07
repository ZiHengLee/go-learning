package dp

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) ==0{
		return 0
	}
	dp := make([][]int,len(text1)+1)
	maxLen := 0
	dp[0] = make([]int,len(text2)+1)
	for k1:=1;k1<=len(text1);k1++ {
		dp[k1] = make([]int,len(text2)+1)
		for k2:=1;k2<=len(text2);k2++{
			if text1[k1-1] == text2[k2-1]{
				dp[k1][k2] = dp[k1-1][k2-1]+1
			}else{
				if dp[k1-1][k2] > dp[k1][k2-1] {
					dp[k1][k2] = dp[k1-1][k2]
				} else {
					dp[k1][k2] = dp[k1][k2-1]
				}
			}
			if dp[k1][k2]>=maxLen{
				maxLen = dp[k1][k2]
			}
		}
	}
	return maxLen
}
