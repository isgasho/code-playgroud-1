package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 解题思路：由于k是非负的，首先要考虑k值比链表还长的情况，其次还有各种边界值，空链表，以及不需要移动的情况
 // 时间复杂度：O(n), 空间复杂度O(1)  beats 双 100% 还是大体满意

func getLenght(head *ListNode) (int, *ListNode) {
	lenght := 0
	var dummy, last *ListNode
	dummy = head
	for dummy != nil {
		lenght += 1
		last = dummy
		dummy = dummy.Next
	}
	return lenght, last
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	length, last := getLenght(head)
	k = k % length
	if k <= 0 {
		return head
	}
	left, dummy := head, head

	for i := 1; i < length - k; i++ {
		left = left.Next
	}
	head = left.Next
	left.Next = nil
	last.Next = dummy
	return head
}
