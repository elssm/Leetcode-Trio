# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def getIntersectionNode(self, headA, headB):
        """
        :type head1, head1: ListNode
        :rtype: ListNode
        """
        res1 = []
        res2 = []
        while headA != None:
            res1.append(headA)
            headA = headA.next
        while headB != None:
            res2.append(headB)
            headB = headB.next
        for i in res1:
            if i in res2:
                return i
        return None
