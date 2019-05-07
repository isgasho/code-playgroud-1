package main

/*
  解题思路：运用和2sum一样的方法，采用3指针。先排序，然后前后逼近求解，貌似有O(n)的算法，但是需要额外空间
  时间复杂度: O(n^2) 空间复杂度: O(1)
 */

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var start, end, sum, closet int
	closet = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		start, end = i+1, len(nums)-1
		for start < end {
			sum = nums[i] + nums[start] + nums[end]
			if sum == target {
				return sum
			} else if sum > target {
				end--
			} else {
				start++
			}
			if abs(sum-target) < abs(closet-target) {
				closet = sum
			}
		}
	}
	return closet
}
