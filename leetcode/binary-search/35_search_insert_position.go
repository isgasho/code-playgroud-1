package main

/*
  解题思路：二分法的变种，需要做的是如果没找到该元素，则返回该元素的插入位置，那么就在二分法结束之后对left进行判断。
  时间复杂度：O(lgn)
 */

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if nums[left] > target || nums[left] == target {
		return left
	}
	return left + 1
}
