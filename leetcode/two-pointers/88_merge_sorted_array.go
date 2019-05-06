package main

/*
  解题思路：分别从两个切片的尾部比较，如果nums2的尾部大，直接加至nums1的末端，注意nums2最后还有剩余的case，需要把剩下的数据移过去
  时间复杂度O(n)
 */

func merge(nums1 []int, m int, nums2 []int, n int) {
	ix1, ix2, end := m-1, n-1, n+m-1
	for ix1 >= 0 && ix2 >= 0 {
		if nums1[ix1] < nums2[ix2] {
			nums1[end] = nums2[ix2]
			ix2--
			end--
		} else {
			nums1[end] = nums1[ix1]
			ix1--
			end--
		}
	}
	for j := 0; j <= ix2; j++ {
		nums1[j] = nums2[j]
	}
}
