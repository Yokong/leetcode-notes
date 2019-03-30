package src

func RomanToInt(s string) int {
	data := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var lastNum int
	var result int
	for i := len(s) - 1; i >= 0; i-- {
		l := data[s[i]]
		sign := 1
		if l < lastNum {
			sign = -1
		}
		result += sign * l
		lastNum = l
	}
	return result
}
