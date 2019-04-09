package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 解题思路：由于两个链表已经是有序的，只要每次比较头节点就行，注意临界条件
 // 时间复杂度：O(n+m) 空间复杂度：O(1)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	curNode := &ListNode{}
	pHead := curNode
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			(*curNode).Next = l2
			l2 = l2.Next
			curNode = curNode.Next
		} else {
			(*curNode).Next = l1
			l1 = l1.Next
			curNode = curNode.Next
		}
	}
	if l1 != nil {
		curNode.Next = l1
	}
	if l2 != nil {
		curNode.Next = l2
	}
	return pHead.Next
}

// TODO 加test