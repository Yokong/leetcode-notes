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