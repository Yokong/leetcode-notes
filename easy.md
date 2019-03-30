### 两数之和

> Go

```
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

```
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