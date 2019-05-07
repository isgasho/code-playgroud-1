package main

/*
  解题思路：双指针，前后进行，滤掉非字母非数字。 坑点在忘记数字了，导致一个"0P"的case看了半天，毕竟在lt里长得像o
  时间复杂度: O(n) 空间复杂度: O(1)
 */

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	l, r := 0, 0
	for left < right {
		l, r = int(s[left]), int(s[right])
		// Delete non-alpha and non-numeric
		if l < 48 || (57 < l && l < 65) || (90 < l && l < 97) || l > 122 {
			left++
			continue
		}
		// Delete non-alpha and non-numeric
		if r < 48 || (57 < r && r < 65) || (90 < r && r < 97) || r > 122 {
			right--
			continue
		}
		// lower
		if l > 91 {
			l -= 32
		}
		if r > 91 {
			r -= 32
		}
		if l != r {
			return false
		} else {
			left++
			right--
			continue
		}
	}
	return true
}
