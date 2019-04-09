package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题思路：使用头插法，需要注意的是 golang 中的 怎么声明 nil 的指针
// 时间复杂度 O(n)

func reverseList(head *ListNode) *ListNode {
	var dummyP *ListNode
	tmpP := &ListNode{}
	for head != nil {
		tmpP = head.Next
		head.Next = dummyP
		dummyP = head
		head = tmpP
	}
	return dummyP
}

// Follow up:
// A linked list can be reversed either iteratively or recursively. Could you implement both?
// 既然提示了可以用迭代和递归解，再来试试递归
