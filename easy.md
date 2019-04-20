### 1. 两数之和
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
### 7. 整数反转
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

### 9. 回文数
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
```

### 14. 最长公共前缀
> 按字母顺序排序数组, 取出最小和最大的字符串的公共前缀就是整个数组的公共前缀
```go
import "sort"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {return ""}
    sort.Strings(strs)
    min := strs[0]
    max := strs[len(strs)-1]
    res := ""
    for i:=0; i<len(min); i++ {
        if min[i] == max[i] {
            res += string(min[i])
        } else {
            break
        }
    }
    return res
}
```

### 20. 有效的括号
> 利用栈结构进行括号配对, 遇到左括号就push到栈中, 遇到右括号从栈顶取出匹配  
> 首先遍历字符串s, 得到每个括号的int32形式数据 `for _, i := range`  
> 如果当前括号为左括号就push数据到stack中  
> 如果当前为右括号, 并且stack为空, 说明没有要匹配的左括号, 直接返回false  
> 如果以上条件都不满足, 就从stack弹出栈顶括号, 来否匹配右括号, 匹配利用map映射{右:左}  
> 上面都不满足, 直接返回false  
> 最后判断stack是否为空, 为空说明所有括号都正确匹配, 返回true, 否则返回false
```go
func isValid(s string) bool {
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
```

### 21. 合并两个有序链表
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    
    var res *ListNode
    if l1.Val >= l2.Val {
        res = l2
        res.Next = mergeTwoLists(l1, l2.Next)
    } else {
        res = l1
        res.Next = mergeTwoLists(l1.Next, l2)
    }
    return res
}
```


### 26. 删除排序数组中的重复项
```go
func removeDuplicates(nums []int) int {
	idx := 0
	for i, v := range nums {
		if i == 0 || v == nums[idx] {
			continue
		}
		if v < nums[idx] {
			break
		}
		nums[idx+1] = v
		idx++
	}
	return idx+1
}
```

### 27. 移除元素
```go
func removeElement(nums []int, val int) int {
    j := 0
    for i, v := range nums {
        if v != val {
            nums[j] = nums[i]
            j++
        }
    }
    return j
}
```

### 28. 实现strStr()
```go
func StrStr(haystack string, needle string) int {
	if needle == "" {return 0}
	if len(needle) > len(haystack) {return -1}
	for i:=0; i<=len(haystack)-len(needle); i++ {
	    if haystack[i:len(needle)+i] == needle {
		    return i
	    }
	}
	return -1
}
```

### 35. 搜索插入位置
> 利用二分查找
```go
func searchInsert(nums []int, target int) int {
    left := 0
    right := len(nums)
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] > target {
            right = mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            return mid
        }
    }
    return left
}
```

### 38. 报数
```go
func CountAndSay(n int) string {
	if n == 1 {return "1"}
	last := CountAndSay(n-1)

	res := ""
	var c byte
	c = last[0]
	count := 1
	for i:=1; i<len(last); i++ {
		if last[i] == c {
			count++
			continue
		}
		res = res + strconv.Itoa(count) + string(c)
		c = last[i]
		count = 1
	}

	res = res + strconv.Itoa(count) + string(c)
	return res
}
```

### 53. 最大子序和
```go
func MaxSubArray(nums []int) int {
	sum := 0
	max := nums[0]
	for i:=0; i<len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
```

### 58. 最后一个单词的长度
```go
func LengthOfLastWord(s string) int {
	l := len(s) - 1
	res := 0
	for l>=0 && s[l] == ' ' {l--}
	for l>=0 && s[l] != ' ' {
		res++
		l--
	}
	return res
}
```

### 66. 加一
```go
func PlusOne(digits []int) []int {
	for i:=len(digits)-1; i>=0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}
```

### 67. 二进制求和
```go
func AddBinary(a, b string) string {
	lenA := len(a)
	lenB := len(b)
	if lenB > lenA {
		a, b = b, a
		lenA, lenB = lenB, lenA
	}
	b = strings.Repeat("0", lenA-lenB) + b
	lenB = len(b)

	carry := byte(0)
	ret := make([]byte, lenB+1)

	for i := lenB - 1; i >= 0; i-- {
		numA := a[i] - '0'
		numB := b[i] - '0'
		sum := numA + numB + carry
		ret[i+1] = sum&1 + '0'
		carry = sum >> 1
	}

	if carry == 0 {
		ret = ret[1:]
	} else {
		ret[0] = carry + '0'
	}
	return string(ret)
}
```

### 69. x 的平方根
```go
func MySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	i := x / 2
	for i*i > x {
		i = (i + x/i) / 2
	}
	return i
}
```