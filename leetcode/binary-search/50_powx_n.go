package main

/*
  解题思路：该题肯定要把时间复杂度降下来。采用递归是比较简单的，但是这个题目在二分法里面，没太明白-，-
  时间复杂度：O(lgn)
 */

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	if n % 2 == 0 {
		return myPow(x*x, n/2)
	} else {
		return myPow(x*x, n/2) * x
	}
}
