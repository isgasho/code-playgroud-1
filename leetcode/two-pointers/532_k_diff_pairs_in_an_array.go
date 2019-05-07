package main

/*
  解题思路：这类问题，考虑用hash存储，可以实现O(n)的算法。坑点在于重复数据只计算一次
  时间复杂度: O(n) 空间复杂度: O(n)
 */

func findPairs(nums []int, k int) int {
	count := 0
	bottle := make(map[int]int)
	for _, i := range nums {
		if k > 0 {
			// i第一次出现，且i-k出现过
			if bottle[i-k] > 0 && bottle[i] == 0 {
				count++
			}
			if bottle[i+k] > 0 && bottle[i] == 0 {
				count++
			}
		}
		bottle[i]++
	}
	if k == 0 {
		for _, v := range bottle {
			if v > 1 {
				count++
			}
		}
	}
	return count
}
