package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 参考：https://leetcode.com/problems/linked-list-cycle-ii/discuss/214353/My-clear-Python-solution.
 // 解题思路：依然用到快慢指针，当第一次快慢指针相遇时。慢指针跑的路程是快指针的一般，那么 2(x1 + x2) = x1 + x2 + x3
 //         推出 x1 = x3。
 // 复杂度: O(n), O(1)

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, quick := head, head
	// 找到相遇节点
	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if quick == slow {
			slow = head
			for slow != quick {
				slow = slow.Next
				quick = quick.Next
			}
			return slow
		}
	}
	return nil
}
