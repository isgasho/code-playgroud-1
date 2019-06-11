## 经典排序算法
1. 排序算法在时间复杂度上分为三个档次：O(n), O(nlgn)，O(n^2)
2. 排序算法的稳定性。如果待排序的列表中存在相同排序值的元素，在排序前后相同排序值的元素排序后相对位置不变。
3. 是否原地排序。也就是说算法是否需要额外空间。

##### 冒泡排序
冒泡比较简单，每次选出子序列的最大值，放置最前端。
时间复杂度: O(n^2)
空间复杂度: O(1)
是否稳定：是
是否原地排序：是

```go
package main

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
			    nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
```