# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def kthLargest(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: int
        """
        res = []

        def inorder(node):
            if node == None:
                return res
            inorder(node.left)
            res.append(node.val)
            inorder(node.right)

        inorder(root)
        res.reverse()
        return res[k - 1]
