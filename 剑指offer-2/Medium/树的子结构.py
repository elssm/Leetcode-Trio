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
        # 如果A和B有一个是空，直接返回False
        if not A or not B:
            return False
        # 判断A的整体结构和B的整体结构或者A的左子树和B的整体或者A的右子树和B的整体
        return self.is_sub(A, B) or self.isSubStructure(A.left, B) or self.isSubStructure(A.right, B)

    def is_sub(self, A, B):
        if not B:
            return True
        if not A or A.val != B.val:
            return False
            # 分别判断左右子树的值是否相等
        return self.is_sub(A.left, B.left) and self.is_sub(A.right, B.right)
