package main

/*
  解题思路：基本是和3sum 是一样的解法，对于这种ksum问题，这种方法只能把时间复杂度降低一个纬度，注意去除相同的case。
  时间复杂度：O(n^3) 空间复杂度：O(1)
 */

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	var left, right int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left = j + 1
			right = len(nums) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
