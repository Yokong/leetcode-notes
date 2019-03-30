### 两数之和

> Go

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

> Python

```python
class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        
        result = {}
        for i, v in enumerate(nums):
            if target - v in result:
                return [result[target - v], i]
            result[v] = i
        return []
```

### 整数反转

> Go

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

> Python

```python
class Solution:
    def reverse(self, x: int) -> int:
        rev = int(str(x)[::-1]) if x > 0 else int(str(-x)[::-1]) * -1
        return rev if -2 ** 31 < rev < 2 ** 31 -1 else 0
```