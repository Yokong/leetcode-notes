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
	bracketMap := map[int32]int32{
		40: 41,
		91: 93,
		123: 125,
	}
	var stack []int32
	for _, i := range s {
		if _, ok := bracketMap[i]; ok {
			stack = append(stack, i)
		} else {
			// 取出上一个
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			bracket := bracketMap[last]
			if i != bracket {
				return false
			}
			continue
		}
	}
	return true
}
