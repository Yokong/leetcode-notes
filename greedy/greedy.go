package greedy

import (
	"sort"
)

func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gg, ss := 0, 0
	for gg < len(g) && ss < len(s) {
		if s[ss] >= g[gg] {
			gg++
		}
		ss++
	}

	return gg
}

// 利用局部最优去推导全局最优
func CanCompleteCircuit(gas []int, cost []int) int {
	sum := 0   // 局部剩余油量和
	total := 0 // 总的剩余油量和
	start := 0 // 起始点

	for i := 0; i < len(cost); i++ {
		sum += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if sum < 0 {
			start = i + 1
			sum = 0
		}
	}

	if total < 0 {
		return -1
	}
	return start
}
