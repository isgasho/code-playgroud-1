package main

/*
  解题思路：一看整数如果用算术来做有点麻烦，简单的方法是转换成字符串来解决，利用前后指针判断字符
  时间复杂度: O(n)
 */

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	st := strconv.Itoa(x)
	strLen := len(st)
	mid := strLen >> 1
	for i := 0; i < mid; i++ {
		if st[i] != st[strLen-i-1] {
			return false
		}
	}
	return true
}
