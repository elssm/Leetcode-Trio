class Solution(object):
    def exist(self, board, word):
        """
        :type board: List[List[str]]
        :type word: str
        :rtype: bool
        """

        def dfs(board, i, j, word, k):
            # 判断下标是否越界
            if i < 0 or i == len(board) or j < 0 or j == len(board[0]):
                return False
            # 如果当前值和word当前值不相等，直接返回
            if board[i][j] != word[k]:
                return False
            # 如果访问到了word最后一个值，返回True
            if k == len(word) - 1:
                return True
            # 先将当前值置空，防止后面重复访问
            board[i][j] = ''
            res = dfs(board, i - 1, j, word, k + 1) or dfs(board, i, j - 1, word, k + 1) \
                  or dfs(board, i + 1, j, word, k + 1) or dfs(board, i, j + 1, word, k + 1)
            board[i][j] = word[k]
            return res

        for i in range(len(board)):
            for j in range(len(board[i])):
                if dfs(board, i, j, word, 0):
                    return True
        return False
