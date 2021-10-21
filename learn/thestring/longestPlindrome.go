package thestring
//动态规划，当下回文串一定由前一个回文串构成
func longestPalindrome(s string) string {
	length := len(s)
	dp := make([][]int, length)
	start, end := 0, 0
	maxLenth := 0

	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = 1
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i-j == 1 {
					dp[j][i] = 2
				} else if dp[j+1][i-1] > 0 {
					dp[j][i] = dp[j+1][i-1] + 2
				}
			}
			if dp[j][i] > maxLenth{
				maxLenth = dp[j][i]
				start = j
				end = i
			}
		}
	}
	return s[start:end+1]
}
