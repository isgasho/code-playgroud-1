package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 解题思路：基本就是翻转链表加了很多边界条件版，找出mn节点，逆转然后拼接。注意头节点的边界值
 // 时间复杂度：O(n), 空间O(1)  写完一定要自己跑几个case

func reverse(head *ListNode) *ListNode {
	var tmpNode, dummy *ListNode
	for head != nil {
		tmpNode = head.Next
		head.Next = dummy
		dummy = head
		head = tmpNode
	}
	return dummy
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	if m > n {
		m, n = n, m
	}

	var nextNode, preNode *ListNode
	left := head
	// 先找翻转的左节点
	for i := 1; i < m; i++ {
		preNode = left
		left = left.Next
	}
	// 非头节点
	if preNode != nil {
		preNode.Next = nil
	}
	// 找到左节点为left，再找右节点right
	right := left
	for i := 0; i < n-m; i++ {
		right = right.Next
		nextNode = right.Next
	}
	// 将右节点中断
	right.Next = nil
	// 翻转中间部分
	if preNode == nil {
		head = reverse(left)
	} else {
		preNode.Next = reverse(left)
	}
	left.Next = nextNode
	return head
}
