package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // 解题思路：找到中点，递归生成搜索树，生成的树是有很多种的，随笔一种就行
 // 时间复杂度：O(nlgn) , 空间复杂度O(1)  应该有更快的解乏，这个只beats 50%

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	var slow, quick *ListNode
	slow, quick = head, head
	var preNode *ListNode
	for quick != nil && quick.Next != nil {
		preNode = slow
		quick = quick.Next.Next
		slow = slow.Next
	}

	next := slow.Next
	slow.Next = nil
	preNode.Next = nil

	root := &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(next),
	}
	return root
}
