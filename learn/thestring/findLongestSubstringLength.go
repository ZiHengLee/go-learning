package thestring
//最长不重复字串
//"abbba"
func lengthOfLongestSubstring(s string) int{
	if len(s) <= 1{
		return len(s)
	}
	helpChar := make([]int, 128)
	for k,_ := range helpChar{
		helpChar[k] = -1
	}
	var maxLenth, curLenth int

	for k,v := range s{
		index := int(v)
		prePosition := helpChar[index]
		distance := k - prePosition
		//后面这个判断语句为了处理中间包含这种情况"abbba"
		if helpChar[index] == -1 || distance > curLenth{
			curLenth += 1
		}else{
			if curLenth > maxLenth {
				maxLenth = curLenth
			}
			curLenth = distance
		}
		helpChar[index] = k
		if curLenth > maxLenth{
			maxLenth =curLenth
		}
	}
	return maxLenth
}
