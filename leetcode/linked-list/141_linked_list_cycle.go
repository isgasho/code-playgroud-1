package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * 解题思路：可以用map存储访问过的节点，但是这样的话会带来空间O(n)的消耗，所以可以用快慢指针的方式来解决
 * 快慢指针的终结点？如果不是循环链表，那么最后一个节点的Next肯定为nil
 * 时间复杂度：O(n) 空间复杂度：O(1)
*/

func hasCycle(head *ListNode) bool {
	slowP := head
	quickP := head
	for slowP != nil && quickP != nil && quickP.Next != nil{
		slowP = slowP.Next
		quickP = quickP.Next.Next
		if quickP == slowP {
			return true
		}
	}
	return false
}