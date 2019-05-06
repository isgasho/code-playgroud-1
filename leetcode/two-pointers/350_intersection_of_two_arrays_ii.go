package main

/*
  解题思路：和题目1类似，但是需要记录数量
  时间复杂度: O(n) 空间复杂度: O(n)

  拓展：
  1. 如果都是有序的切片，怎么优化？有序的话，可以使空间复杂度降为O(1)
  2. 如果nums1比nums2短，怎么优化？
  3. 如果nums2存在硬盘，内存有限制不能一次加载完所有元素，怎么优化？
 */

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	bottle := make(map[int]int)
	for _, i := range nums1 {
		bottle[i]++
	}
	for _, j := range nums2 {
		if count, ok := bottle[j]; ok && count > 0 {
			res = append(res, j)
			bottle[j]--
		}
	}
	return res
}
