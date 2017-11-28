package main
func RemoveEven(sl []int) []int {
	var ans []int
	for _, v := range sl {
		if v % 2 != 0 {
		ans = append(ans, v)
		}
	}
	return ans
}

func PowerGenerator(a int) func() int {
	sum := a
	return func() int {
		sum *= a
		return sum
	}
}

func DifferentWordsCount(str string) int {
	str += string('1')
	var str1 string
	m := make(map[string]int)
	for _, v := range str {
		vtmp := rune(v)
		if unicode.IsLetter(vtmp) {
		vv := strings.ToLower(string(v))
		str1 += vv
		
		} else {
		if str1 != "" {
			m[str1] += 1
			str1 = ""
		}
		}
	}
	return len(m)
}
