package main

/* 解题思路：题目条件给出必有解，可以采用双指针，左右相加，如果和大于目标值，则需要降低，所以右指针后退。同理
   note: 二分也可以实现一个O(nlgn)的解法，不过不是最优解
   时间复杂度：O(n) 空间复杂度O(n)
*/

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	sum := numbers[left] + numbers[right]
	for sum != target {
		if sum > target {
			right--
		} else {
			left++
		}
		sum = numbers[left] + numbers[right]
	}
	return []int{left + 1, right + 1}
}
