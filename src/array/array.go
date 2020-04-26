package array

import (
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
