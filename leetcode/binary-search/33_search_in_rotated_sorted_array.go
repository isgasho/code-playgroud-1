package main

/*
  解题思路：看到这个题，做的时候没有想到偏移的规律，有点难解。那么如何解呢？先找到最小值的索引，然后其实这个就是偏移量。那么我们可以用这个偏移量去得到真实的mid索引。之后就跟普通二分查找一样了。
  时间复杂度: O(lgn)
 */

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// Got the smallest value's index is left
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		realMid := (mid + left) % len(nums)
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
