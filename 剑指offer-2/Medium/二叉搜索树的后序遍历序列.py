class Solution(object):
    def verifyPostorder(self, postorder):
        """
        :type postorder: List[int]
        :rtype: bool
        """
        if len(postorder) == 1 or (not postorder):
            return True
        root = postorder[-1]  # 后序遍历最后一个值为根结点
        for i in range(len(postorder)):  # 找到根节点右子树的第一个值
            if postorder[i] >= root:
                break
        left = postorder[:i]  # 划分左子树
        right = postorder[i:len(postorder) - 1]  # 划分右子树
        for i in left:  # 如果左子树中有小于根的，直接return
            if i > root:
                return False
        for i in right:  # 如果右子树中有小于根的，直接return
            if i < root:
                return False
        left = self.verifyPostorder(left)  # 递归判断左子树
        right = self.verifyPostorder(right)  # 递归判断右子树
        if left and right:
            return True
        return False
