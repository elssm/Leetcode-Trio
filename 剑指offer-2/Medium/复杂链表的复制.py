"""
# Definition for a Node.
class Node:
    def __init__(self, x, next=None, random=None):
        self.val = int(x)
        self.next = next
        self.random = random
"""


class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: Node
        :rtype: Node
        """
        if head == None:
            return None
        d = {}
        node = head  # 保留头节点不改变
        while (node):  # 先用字典单独复制每一个节点信息
            d[node] = Node(node.val)
            node = node.next
        node = head  # 保留头节点不改变
        while (node):  # 链接每一个单节点
            d[node].next = d.get(node.next)
            d[node].random = d.get(node.random)
            node = node.next
        return d[head]
