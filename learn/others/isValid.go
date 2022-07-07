package others

//有效括号
func isValid(s string) bool {
	help := make([]int32, 0)
	for _, v := range s {
		if len(help) == 0 {
			help = append(help, v)
		} else {
			lastIndex := len(help) - 1
			if (v == 41 && help[lastIndex] == 40) || (v == 125 && help[lastIndex] == 123) || (v == 93 && help[lastIndex] == 91) {
				help = help[:lastIndex]
			} else {
				help = append(help, v)
			}
		}

	}
	if len(help) == 0 {
		return true
	}
	return false
}
