# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if root == None:
            return True
        left = self.hight(root.left)
        right = self.hight(root.right)
        if abs(left - right) > 1:
            return False
        else:
            return self.isBalanced(root.left) and self.isBalanced(root.right)

    def hight(self, root):
        if root == None:
            return 0
        return max(self.hight(root.left), self.hight(root.right)) + 1
