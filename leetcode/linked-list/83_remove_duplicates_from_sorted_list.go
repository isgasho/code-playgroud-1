package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 解题思路：由于是有序的链表，那么重复一定是连续的，所以记录一个preNode节点就可以解决
 // 时间复杂度: O(n) 空间复杂度：O(1)

func deleteDuplicates(head *ListNode) *ListNode {
	var preNode, dummyP *ListNode
	dummyP = head
	for head != nil {
		// 元素重复 删除该点 跳下一个
		if preNode != nil && preNode.Val == head.Val {
			preNode.Next = head.Next
			head = head.Next
			continue
		}
		preNode = head
		head = head.Next
	}
	return dummyP
}
