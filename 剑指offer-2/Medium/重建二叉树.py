# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def buildTree(self, preorder, inorder):
        """
        :type preorder: List[int]
        :type inorder: List[int]
        :rtype: TreeNode
        """
        if not preorder:
            return None
        root = TreeNode(preorder[0])  # 将前序遍历第一个值作为根节点
        index = inorder.index(preorder[0])  # 找到根节点再中序遍历中的位置
        root.left = self.buildTree(preorder[1:1 + index], inorder[:index])  # 递归构建左子树
        root.right = self.buildTree(preorder[1 + index:], inorder[index + 1:])  # 递归构建右子树
        return root
