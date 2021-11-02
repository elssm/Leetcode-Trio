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
        if head == None or head.next == None:
            return head
        node = None
        curnode = head
        while curnode:
            temp = curnode.next
            curnode.next = node
            node = curnode
            curnode = temp
        return node
