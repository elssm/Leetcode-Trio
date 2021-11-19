class Solution(object):
    def movingCount(self, m, n, k):
        """
        :type m: int
        :type n: int
        :type k: int
        :rtype: int
        """

        def dfs(i, j, k):
            if not 0 <= i < m or not 0 <= j < n or (i, j) in a or (i % 10 + i // 10 + j % 10 + j // 10) > k: return 0
            # 把这个点加入到访问集合中
            a.add((i, j))
            # 因为从（0,0）开始，所以只要向下走或向右走就行
            dfs(i + 1, j, k)
            dfs(i, j + 1, k)

        a = set()
        dfs(0, 0, k)
        return len(a)
