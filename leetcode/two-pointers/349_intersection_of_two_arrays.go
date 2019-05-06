package main

/*
  解题思路：需要借助辅助空间，可以实现O(n)的时间复杂度
  时间复杂度: O(n) 空间复杂度：O(n)
 */

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	bottle := make(map[int]bool)
	for _, i := range nums1 {
		bottle[i] = false
	}
	for _, i := range nums2 {
		v, ok := bottle[i]
		if ok && v == false {
			bottle[i] = true
			res = append(res, i)
		}
	}
	return res
}
