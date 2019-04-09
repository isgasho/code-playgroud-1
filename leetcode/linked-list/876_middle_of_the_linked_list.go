package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题思路：依然可以用快慢指针，快指针跑两步，慢指针跑一步
// 时间复杂度：O(n)

func middleNode(head *ListNode) *ListNode {
	slowP := head
	quickP := head
	for quickP != nil && quickP.Next != nil {
		quickP = quickP.Next.Next
		slowP = slowP.Next
	}
	return slowP
}
