package main

/*
  解题思路：题是一道好题，但是描述真的一言难尽，所以有3000多个反对票啦。当然这道题还是典型的双指针，一个指针指向不同元素的最后，一个指向遍历，然后交换，最后返回第一个指针的索引加1就行。
  时间复杂度: O(n) 空间复杂度: O(1)
 */

func removeDuplicates(nums []int) int {
	var i int = 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}
