package dp

//最短编辑距离

//dp[i][j] word1的前i位到word2的前j位所需要的最短步数
func minDistance(word1 string, word2 string) int {
	len1,len2 := len(word1)+1,len(word2)+1
	dp := make([][]int,len1)
	for j := 0;j < len1; j++{
		tmp := make([]int,len2)
		dp[j] = tmp
		dp[j][0] = j
		for k := 0; k< len2;k++{
			dp[0][k] = k
		}
	}
	for i:=1;i<len1;i++{
		for j:=1;j<len2;j++{
			if word1[i-1]==word2[j-1]{
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] =minFunc(minFunc(dp[i-1][j-1],dp[i-1][j]),dp[i][j-1])+1
			}
		}
	}
	return dp[len1-1][len2-1]
}

func minFunc(a,b int) int{
	if a< b{
		return a
	}
	return b
}
