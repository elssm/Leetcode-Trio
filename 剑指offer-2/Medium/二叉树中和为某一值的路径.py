# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def pathSum(self, root, target):
        """
        :type root: TreeNode
        :type target: int
        :rtype: List[List[int]]
        """
        res = []
        result = []

        def dfs(root, tmp):
            if root == None:
                # 根为空，直接返回
                return
            if root.left == None and root.right == None:
                # 已经遍历到了叶子节点，加入叶子节点的值后返回
                res.append(tmp + [root.val])
                return
            dfs(root.left, tmp + [root.val])
            dfs(root.right, tmp + [root.val])

        dfs(root, [])

        # 遍历res选出和为targetSum的存入result中
        for i in res:
            if sum(i) == target:
                result.append(i)
        return result
