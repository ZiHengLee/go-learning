package thestring

//全排列
func Permutation(str *string,resList *[]string,l,r int){
	strr := []rune(*str)
	if l == r-1 {
		*resList = append(*resList, string(strr))
	}else{
		tempSet := map[rune]interface{}{}
		tempSet[strr[l]] = nil
		for j := l;j<r;j++{
			if j == l{
				*resList = append(*resList, string(strr))
			}

			if _,ok := tempSet[strr[j]];ok{
				continue
			}else{
				tempSet[strr[j]] = nil
			}
			strr[l],strr[j] = strr[j],strr[l]
			tempStr := string(strr)
			Permutation(&tempStr,resList,l+1,r)
		}
	}
}

func PermutationInt(arrys []int,resList *[][]int,l,r int){
	if l == r {
		copyData := make([]int, len(arrys))
		copy(copyData, arrys)
		*resList = append(*resList, copyData)
		return
	}
	tempSet := map[int]interface{}{}
	for j := l;j<r;j++{
		if _,ok := tempSet[arrys[j]];ok{
			continue
		}else{
			tempSet[arrys[j]] = nil
		}
		arrys[l],arrys[j] = arrys[j],arrys[l]
		PermutationInt(arrys,resList,l+1,r)
		arrys[l],arrys[j] = arrys[j],arrys[l]
	}
}

//全排列2 https://leetcode.cn/problems/permutations-ii/
//func permuteUnique(nums []int) [][]int {
//
//}