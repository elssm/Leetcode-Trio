# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def isSubStructure(self, A, B):
        """
        :type A: TreeNode
        :type B: TreeNode
        :rtype: bool
        """
        if not A or not B:
            return False
        return self.helper(A, B) or self.isSubStructure(A.left, B) or self.isSubStructure(A.right, B)

    def helper(self, A, B):
        if not B:
            return True
        if not A or A.val != B.val:
            return False
            # 分别判断左右子树的值是否相等
        return self.helper(A.left, B.left) and self.helper(A.right, B.right)
