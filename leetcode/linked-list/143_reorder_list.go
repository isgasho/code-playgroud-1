package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 解题思路：依然是快慢指针将后半部逆序然后重组，但是有几个注意点：一个是奇偶性，一个是前半部分链表的最后一个节点必须置空

 // 值得重做一遍，O(n), O(1)


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

func reorderList(head *ListNode) {
	dummyP, midP := head, head
	tmp := &ListNode{}
	if head == nil || head.Next == nil {
		return
	}
	for dummyP != nil && dummyP.Next != nil {
		tmp = midP
		dummyP = dummyP.Next.Next
		midP = midP.Next
	}
	tmp.Next = nil
	midP = reverse(midP)
	dummyP = head
	var tmpNode *ListNode
	for dummyP.Next != nil {
		tmpNode = dummyP.Next
		dummyP.Next = midP
		midP = midP.Next
		dummyP.Next.Next = tmpNode
		dummyP = dummyP.Next.Next
	}
	dummyP.Next = midP
}
