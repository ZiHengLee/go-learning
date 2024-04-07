package utils

func Swap(str *string,left,right int){
	strr := []rune(*str)
	strr[left],strr[right] = strr[right],strr[left]
	*str = string(strr)
}
