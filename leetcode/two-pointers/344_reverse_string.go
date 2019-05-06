package main

/*
  解题思路：题目限制O(1)的时间复杂度，那么只能左右指针对调
  时间复杂度: O(n)，空间复杂度O(1)
 */

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
