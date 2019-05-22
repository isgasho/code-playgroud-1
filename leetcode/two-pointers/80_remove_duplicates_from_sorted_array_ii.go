package main

/*
  解题思路：想了几个方法，但是都卡在临界点。于是看了一个java的解答，没有注释没看懂，跑了一遍case才理解
           又简洁又好懂。其实就是像滑动窗口的思想，滑动窗口为3，以 [0,0,1,1,2,2,3,3] 为例子。可以发现规律是窗口末尾始终要大于窗口头元素。
           回过头再把 第一个题改了
  时间复杂度：O(n) 空间复杂度: O(1)
 */

func removeDuplicates(nums []int) int {
	var i int = 0
	for _, v := range nums {
		if i < 2 || v > nums[i-2] {
			nums[i] = v
			i++
		}
	}
	return i
}
