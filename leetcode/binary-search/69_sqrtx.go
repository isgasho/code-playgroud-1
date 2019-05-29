package main

/*
  解题思路：看题描述，基本是一个二分查找。如果不用二分得挨个计算，需要对最后结果做一下校准。
  时间复杂度: O(lgn)
 */

func mySqrt(x int) int {
	left, right := 1, x/2
	for left < right {
		mid := (left + right) >> 1
		if mid * mid > x {
			right = mid - 1
		} else if mid * mid == x {
			return mid
		} else {
			left = mid + 1
		}
	}
	if left * left > x {
		return left - 1
	}
	return left
}
