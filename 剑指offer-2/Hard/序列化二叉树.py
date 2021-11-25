# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Codec:

    def __init__(self):
        self.strTree = []

    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        if not root:
            self.strTree.append("null")
            return self.strTree
        curStr = str(root.val)
        self.strTree.append(curStr)
        self.serialize(root.left)
        self.serialize(root.right)
        return self.strTree

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        if len(data) == 1 and data[0] == "null":
            return None
        Tval = data.pop(0)
        if Tval == "null":
            return None
        root = TreeNode(Tval)
        root.left = self.deserialize(data)
        root.right = self.deserialize(data)
        return root


# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
