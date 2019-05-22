package main

/*
  解题思路：双指针，一个指针指向小于x 一个大于x 最后合并
  时间复杂度：O(n)  空间复杂度：O(1)
 */

func partition(head *ListNode, x int) *ListNode {
	first, second := &ListNode{}, &ListNode{}
	dummyFirst, dummySecond := first, second
	for head != nil {
		if head.Val < x {
			fmt.Println(head.Val)
			first.Next = head
			first = first.Next
		} else {
			second.Next = head
			second = second.Next
		}
		head = head.Next
	}
	second.Next = nil
	first.Next = dummySecond.Next
	return dummyFirst.Next
}
