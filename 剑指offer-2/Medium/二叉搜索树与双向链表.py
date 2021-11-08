"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
"""


class Solution(object):
    head = None
    tail = None

    def treeToDoublyList(self, root):
        """
        :type root: Node
        :rtype: Node
        """
        if root == None:
            return None
        self.pre(root)  # 中序遍历
        self.head.left = self.tail  # 头节点前驱指向尾节点
        self.tail.right = self.head  # 尾节点后驱指向头节点
        return self.head

    def pre(self, root):
        if root == None:
            return
        self.pre(root.left)
        if (self.head == None):
            self.head = self.tail = root
        else:
            self.tail.right = root
            root.left = self.tail
            self.tail = root
        self.pre(root.right)
