package others

//z字形反转
func convert(s string, numRows int) string {
	if numRows == 1{
		return s
	}
	sLen := len(s)
	dived := sLen % (2*numRows - 2)
	div := sLen / (2*numRows - 2)
	if dived >= 1 && dived <= numRows {
		div = 1 + div*(numRows-1)
	} else {
		div += 1 + (dived % numRows) + div*(numRows-1)
	}
	newArray := make([][]byte, numRows)
	for i, _ := range newArray {
		newArray[i] = make([]byte, div)
	}
	//fmt.Println(len(newArray),len(newArray[0]))
	i := 0
	j := 0
	direction := -1
	for start := 0; start < sLen; start++ {
		newArray[j][i] = s[start]
		//fmt.Println(string(newArray[j][i]), j, i,start)
		if j == numRows-1{
			direction = 1
		}else if j == 0{
			direction = -1
		}
		if direction == 1{
			i+=1
			j-=1
		}else{
			j+=1
		}
	}
	res := ""
	for _, v := range newArray{
		for i,_ := range v{
			if v[i] == 0{
				continue
			}
			res+=string(v[i])
		}
	}
	return res
}
