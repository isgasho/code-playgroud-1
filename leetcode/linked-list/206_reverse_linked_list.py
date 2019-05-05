# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        rev_head = None
        while head is not None:
            tmp = head.next
            head.next = rev_head
            rev_head = head
            head = tmp
        return rev_head
     
    def reverseList(self, head, last = None):
    	if not head:
            return last
    	next = head.next
    	head.next = last
    	return self.reverseList(next, head)
