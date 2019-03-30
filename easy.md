### 两数之和
```golang
func twoSum(nums []int, target int) []int {
    m := make(map[int] int, len(nums))
    
    for i, v := range nums {
        if j, ok := m[target-v]; ok {
            return []int{j, i}
        }
        m[v] = i
    }
    return []int{}
}
```
### 整数反转
```go
import "math"

func reverse(x int) int {
    rev := 0
    MAX := math.MaxInt32 / 10
    MIN := math.MinInt32 / 10
    for x != 0 {
        pop := x % 10
        x /= 10
        if rev > MAX || (rev == MAX && pop > 7) {
            return 0
        }
        if rev < MIN || (rev == MIN && pop < -8) {
            return 0
        }
        rev = rev * 10 + pop
    }
    return rev
}
```

### 回文数
```go
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if x <= 9 {
        return true
    } else if x % 10 == 0 {
        return false
    }
    
    num := 0
    for x > 0 {
        pop := x % 10
        x /= 10
        num = num * 10 + pop
        if x == num || x / 10 == num {
            return true
        }
    }
    return false
}
```

### 13. 罗马数字转整数
```go
func romanToInt(s string) int {
    data := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }    
    
    result := 0
    lastNum := 0
    for i:=len(s) -1; i>=0; i-- {
        cur := data[s[i]]
        sign := 1
        if cur < lastNum {
            sign = -1
        }
        result += cur * sign
        lastNum = cur
    }
    return result
}
```