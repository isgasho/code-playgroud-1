package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 解题思路：这个题中有要求，不能逆置原链表，当然逆置链表还有很多细节要处理。那么要么就用栈存起来计算，想到栈的话，那么其实递归能解决这个问题。
 // 时间复杂度：O(n)，O(1)

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func helper(l1 *ListNode, l2 *ListNode, diff int) (*ListNode, int) {
	if l1 == nil || l2 == nil {
		return nil, 0
	}
	var next *ListNode
	var sum, carry int
	// diff 大于0 l1右移 将l1 l2对齐，继续递归
	if diff > 0 {
		next, carry = helper(l1.Next, l2, diff-1)
		sum = carry + l1.Val
	} else {
		next, carry = helper(l1.Next, l2.Next, diff)
		sum = carry + l1.Val + l2.Val
	}
	return &ListNode{Val: sum % 10, Next: next}, sum / 10
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Len, l2Len := getLength(l1), getLength(l2)
	if l1Len < l2Len {
		l1, l2 = l2, l1
		l1Len, l2Len = l2Len, l1Len
	}
	next, carry := helper(l1, l2, l1Len-l2Len)
	if carry > 0 {
		next = &ListNode{Val: carry, Next: next}
	}
	return next
}
