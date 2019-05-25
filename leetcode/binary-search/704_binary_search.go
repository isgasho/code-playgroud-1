package main

/*
  解题思路：二分法是比较重要的，但是有很多细节需要注意，多写写。
  时间复杂度: O(lgn)
 */

func search(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
