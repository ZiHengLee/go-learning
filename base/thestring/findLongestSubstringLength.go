package thestring

//https://leetcode.cn/problems/longest-substring-without-repeating-characters/
//最长不重复字串
//"abbba"
//直接便利记录每个不同字母的位置，每遇到相同的字母时，计算一次长度，并对比最大长度
//关键点：
//1.遇到相同字母后的启动长度，就是这两段的长度开始重新计算
//2."abbba"
func lengthOfLongestSubstring(s string) int {
	if len(s)<=1{
		return len(s)
	}
	indexStr := make([]int,128)
	for k := range indexStr{
		indexStr[k] = -1
	}
	var maxLen, curLen, tempDis int
	for k,c := range s{
		cNum := int(c)
		prePosition := indexStr[cNum]
		tempDis = k-prePosition
		if prePosition == -1 || tempDis > curLen{
			//后面这个条件为了解决中间已经出现过相同字母的这种情况
			curLen+=1
		}else{
			if curLen > maxLen{
				maxLen = curLen
			}
			curLen = tempDis
		}
		indexStr[cNum] = k
		if curLen > maxLen{
			maxLen=curLen
		}
	}
	return curLen
}
