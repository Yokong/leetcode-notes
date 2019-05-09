package src

import (
	"strconv"
	"strings"
)

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
	if len(s)%2 == 1 {
		return false
	}
	bracketMap := map[int32]int32{
		41:  40,
		93:  91,
		125: 123,
	}
	var stack []int32
	for _, i := range s {
		if i == 40 || i == 91 || i == 123 {
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

// 28. 实现strStr()
func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

// 38. 报数
func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	last := CountAndSay(n - 1)

	res := ""
	var c byte
	c = last[0]
	count := 1
	for i := 1; i < len(last); i++ {
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

// 53. 最大子序和
func MaxSubArray(nums []int) int {
	sum := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
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

// 58. 最后一个单词的长度
func LengthOfLastWord(s string) int {
	l := len(s) - 1
	res := 0
	for l >= 0 && s[l] == ' ' {
		l--
	}
	for l >= 0 && s[l] != ' ' {
		res++
		l--
	}
	return res
}

// 66. 加一
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}

// 67. 二进制求和
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

// 69. x 的平方根
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

// 70. 爬楼梯
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 83. 删除排序链表中的重复元素
func DeleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 88. 合并两个有序数组
func Merge(nums1, nums2 []int, m, n int) {
	i := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
			i--
		} else {
			nums1[i] = nums2[n]
			n--
			i--
		}
	}
	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
}

// 100. 相同的树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
