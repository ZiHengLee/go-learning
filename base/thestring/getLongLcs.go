package thestring

//求两个字符串的最长公共子串，动态规划，当前的公共字串一定由前一组公共字串构成
//以str1第i位字符结尾和str2第j位字符结尾的公共子串的长度，要求最长子串由最后一位结尾
func getLcs(str1,str2 string) (maxLen int) {
	len1,len2 := len(str1),len(str2)
	if len1 == 0 || len2 ==0{
		return
	}
	dp := make([][]int,len1+1)
	for i:=1;i<=len1;i++{
		dp[i] = make([]int,len2+1)
		for j:=1;j<=len2;j++{
			if str1[i-1] == str2[j-1]{
				dp[i][j] = dp[i-1][j-1]+1
			}
			if dp[i][j]>maxLen{
				maxLen=dp[i][j]
			}
		}
	}
	return
}