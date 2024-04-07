package dp

//最长有效括号
// dp的方法有点抽象
//dp[i]为第i位结尾的有效长度
//
//s[i] = ( dp[i]=0
//s[i] = )
//	s[i-1]=( dp[i]=dp[i-2]+2
//	s[i-1]=)
//		首先要夸过这个dp的长度，才能找到对应的下标
//		对应s[i-1]的左边的那个数s[i-dp[i-1]] 再左边是的s[i-1-dp[i-1]]
//		if s[i-1-dp[i-1]] == ( 和 s[i] 对应：
//			dp[i] = 2 + dp[i-1]+ 再前面的dp[i-dp[i-1]-2]

func longestValidParentheses(s string) int {
	sLen := len(s)
	if sLen == 0 || sLen == 1{
		return 0
	}
	if sLen == 2 {
		if s[0] == '(' && s[1] == ')'{
			return 2
		}else{
			return 0
		}
	}
	dp := make([]int, sLen)
	maxLen := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '('{
			dp[i] = 0
		}else if s[i] ==')'{
			if s[i-1] =='('{
				if i >=2 {
					dp[i] = dp[i-2]+2
				}else{
					dp[i] = 2
				}
			}else if s[i-1] ==')' && dp[i-1] > 0{
				if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '('{
					if i-dp[i-1]-2 >=0 && dp[i-dp[i-1]-2] > 0{
						dp[i] = dp[i-1]+2+dp[i-dp[i-1]-2]
					}else{
						dp[i] = dp[i-1]+2
					}
				}
			}
		}
		if dp[i] > maxLen{
			maxLen = dp[i]
		}
	}
	return maxLen
}
