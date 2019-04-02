package src

func RomanToInt(s string) int {
	data := map[int32]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result, last int
	for _, v := range s {
		cur := data[v]
		result += cur
		if cur > last {
			result -= 2 * last
		}
		last = data[v]
	}
	return result
}


// 20. 有效的括号
func IsValid(s string) bool {
	if len(s) % 2 == 1 {return false}
	bracketMap := map[int32]int32{
		41: 40,
		93: 91,
		125: 123,
	}
	var stack []int32
	for _, i := range s {
		if i == 40 || i == 91 || i == 123{
			stack = append(stack, i)
		} else if len(stack) == 0 {
			return false
		} else if stack[len(stack)-1] == bracketMap[i] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
