package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 解题思路：利用快慢指针找到中点，逆转后半部分，然后比较。
 //         关键点在于奇偶性的判断，这个取决于快指针是否为nil。如果是偶数则快慢指针同时到达，快指针为nil
 // 时间空间复杂度由于题目有明确要求，时间复杂度O(n), 空间复杂度O(1)

func reverse(head *ListNode) *ListNode {
	var dummyP, tmpNode *ListNode
	for head != nil {
		tmpNode = head.Next
		head.Next = dummyP
		dummyP = head
		head = tmpNode
	}
	return dummyP
}

func isPalindrome(head *ListNode) bool {
	dummyP, midP := head, head
	for dummyP != nil && dummyP.Next != nil {
		dummyP = dummyP.Next.Next
		midP = midP.Next
	}
	// is odd
	if dummyP != nil {
		midP = midP.Next
	}
	midP = reverse(midP)
	for midP != nil {
		if head.Val == midP.Val {
			head, midP = head.Next, midP.Next
			continue
		}
		return false
	}
	return true
}
