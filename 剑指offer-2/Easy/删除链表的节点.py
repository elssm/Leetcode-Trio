# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution(object):
    def deleteNode(self, head, val):
        """
        :type head: ListNode
        :type val: int
        :rtype: ListNode
        """
        if head == None:
            return None
        if head.val == val:
            head = head.next
            return head
        pre = pre_next = head
        while pre_next != None:
            if pre_next.val == val:
                pre.next = pre_next.next
            pre = pre_next
            pre_next = pre_next.next
        return head
