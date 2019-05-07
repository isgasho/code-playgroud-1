package main

/*
  解题思路：暴力解法可以实现，但是是O(n^2)的时间复杂度，这里仍然可以用前后放大逼近的思路，计算最值
  时间复杂度：O(n) 空间复杂度: O(1)
 */

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func caculate_container(height []int, start, end int) int {
	return min(height[start], height[end]) * (end - start)
}

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	max_container := caculate_container(height, start, end)
	for start < end {
		container := caculate_container(height, start, end)
		if container > max_container {
			max_container = container
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return max_container
}
