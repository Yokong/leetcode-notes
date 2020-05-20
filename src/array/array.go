package array

import (
	"leetcode-notes/src/util"
	"math"
	"sort"
)

// 15 三数之和
// 排序 + 双指针
// 算法流程:
// 		1. 特判，对于数组长度 n，如果数组为 null或者数组长度小于 3，返回 []。
// 		2. 对数组进行排序。
//		3. 遍历排序后数组：
//			3.1 若nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
//			3.2 对于重复元素：跳过，避免出现重复解
// 			3.3 令左指针 L=i+1，右指针 R=n-1，当 L<R 时，执行循环：
//				3.3.1 当 nums[i]+nums[L]+nums[R]==0, 执行循环,
//				判断左界和右界是否和下一位置重复，去除重复解。
//				并同时将 L,R 移到下一位置，寻找新的解
//				3.3.2 若和大于 0，说明 nums[R] 太大，R 左移
// 				3.3.3 若和小于 0，说明 nums[L] 太小，L 右移
//
//复杂度分析
//* 时间复杂度: O(n^2), 数组排序O(N log N), 遍历数组O(n), 双指针遍历O(n),
//总体O(N log N) + O(n) * O(n), O(n^2)
//* 空间复杂度: O(1)
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	var res [][]int
	for i, v := range nums {

		// 3.1
		if v > 0 {
			return res
		}

		// 3.2
		if i > 0 && v == nums[i-1] {
			continue
		}

		// 3.3
		L := i + 1
		R := len(nums) - 1
		for L < R {
			if v+nums[L]+nums[R] == 0 {
				res = append(res, append([]int{}, v, nums[L], nums[R]))
				for L < R && nums[L] == nums[L+1] {
					L += 1
				}
				for L < R && nums[R] == nums[R-1] {
					R -= 1
				}
				L += 1
				R -= 1
			} else if v+nums[L]+nums[R] > 0 {
				R -= 1
			} else if v+nums[L]+nums[R] < 0 {
				L += 1
			}
		}
	}
	return res
}

// 18 四数之和
// 简化四数为三数
// 简化三数为两数
// 算法流程:
//   判断边界
//   排序数组
//   从0开始遍历数组, 设为z, 去除重复数据判断
//   把target-当前值将四数之和退化成三数之和
//   然后从z+1位开始遍历后面的数, 设为k, 同样去除重复数据
//   设置两个指针i, j分别是k+1, 和数组长度-1
//   把target值接着减去上面退化的target, 将三数之和退化成两数之和
//   通过i, j两个指针循环计算满足条件的值
func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)
	length := len(nums)
	var res [][]int

	for z := 0; z < length; z++ {
		if z > 0 && nums[z] == nums[z-1] {
			continue
		}

		newTarget := target - nums[z]

		for k := z + 1; k < length; k++ {
			if k > z+1 && nums[k] == nums[k-1] {
				continue
			}
			i, j := k+1, length-1
			newTarget2 := newTarget - nums[k]
			for i < j {
				if nums[i]+nums[j] == newTarget2 {
					res = append(res, append([]int{}, nums[z], nums[k], nums[i], nums[j]))
					for i < j && nums[i] == nums[i+1] {
						i++
					}
					for i < j && nums[j] == nums[j-1] {
						j--
					}
					i++
					j--
				} else if nums[i]+nums[j] < newTarget2 {
					i++
				} else {
					j--
				}
			}
		}
	}

	return res
}

// 11. 盛最多水的容器
// 双指针
// 过程
//   定义两个指针分别指向数组第一个元素和最后一个元素
//   取两个元素(这里定义为a, b)的最小值乘上宽度: min(a, b) * w w为剩余的宽度
//   将上面得到的结果与定义的res取最大值
//   如果左边的元素小于右边的，就将左指针右移否则就将右指针向左移
//   循环上面步骤直到左右指针相遇为止
func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0

	for l < r {
		res = util.MaxInt(res, util.MinInt(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

//16. 最接近的三数之和
//排序 + 双指针
//	首先进行数组排序, 时间复杂度 O(nlogn)
//	遍历数组nums, 利用下标i形成一个固定值nums[i]
//	定义两个指针l, r分别是i+1, 和数组长度-1
//	计算sum = nums[i] + nums[l] + nums[r]的结果, 判断sum与target的距离, 更新最近那个结果
//	判断sum与target的大小关系, 因为数组有序, 如果sum > target 则r--, 如果sum < target则l++
//		如果sum == target直接返回结果
//	整个遍历过程, 最外for循环n次, 双指针为n次,时间复杂度为O(n^2)
//	总时间复杂度: O(nlogn) + O(n^2) = O(n^2)
func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)

	length := len(nums)
	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < length; i++ {
		l, r := i+1, length-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
				res = sum
			}

			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				return res
			}
		}
	}

	return res
}

// 从后向前查找第一个相邻升序的元素(i,j),满足nums[i] < nums[j]
// 在[j,end)从后向前查找第一个满足nums[i] < nums[k]的k
// 交换nums[i], nums[k]
// 此时[j,end)必然是降序，逆转[j, end),全其升序
func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 二分查找
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
// 用二分查找, 找到与target相同值的位置, 然后判断相邻位置是否相等
func SearchRange(nums []int, target int) []int {
	binarySearch := func(nums []int, target int, left bool) int {
		l, r := 0, len(nums)-1
		for l <= r {
			mid := l + (r-l)/2
			if nums[mid] > target || (left && target == nums[mid]) {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return l
	}

	leftIdx := binarySearch(nums, target, true)
	if leftIdx == len(nums) || nums[leftIdx] != target {
		return []int{-1, -1}
	}

	return []int{leftIdx, binarySearch(nums, target, false) - 1}
}
