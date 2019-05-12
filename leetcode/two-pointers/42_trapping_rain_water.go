package main

/*
  解题思路：这道题之前被考过，当时没有想到怎么做，这里还是典型的双指针，算出每个点左边最高和右边最高然后减去高度即该点的容量。该解法还有很多可优化空间，同时也有空间复杂度O(1)的算法。
  时间复杂度: O(n) 空间复杂度: O(n)
 */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	max_left := []int{}
	var left, right, trapped, n int = 0, 0, 0, 0
	for i := 0; i < len(height); i++ {
		max_left = append(max_left, left)
		if height[i] > left {
			left = height[i]
		}
	}
	max_right := []int{}
	for j := len(height) - 1; j >= 0; j-- {
		max_right = append(max_right, right)
		if height[j] > right {
			right = height[j]
		}
	}
	for i := 0; i < len(height); i++ {
		n = min(max_right[len(max_right)-1-i], max_left[i]) - height[i]
		if n > 0 {
			trapped += n
		}
	}
	return trapped
}
