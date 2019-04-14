package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 解题思路：关键在于合并后的各种临界条件，注意l1 l2的指针移动，写完一定要跑一个case检查

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, val := 0, 0
	head := &ListNode{}
	var dummyP *ListNode = head
	for l1 != nil || l2 != nil || carry == 1 {
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		val += carry
		carry = val / 10
		val = val % 10

		head.Next = &ListNode{Val: val}
		val = 0
		head = head.Next
	}
	return dummyP.Next
}
