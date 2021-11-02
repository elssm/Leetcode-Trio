添加备注


class Solution(object):
    def lastRemaining(self, n, m):
        """
        :type n: int
        :type m: int
        :rtype: int
        """
        res = []
        for i in range(n):
            res.append(i)
        # print(res)
        index = 0
        while n > 1:
            index = (index + m - 1) % n
            # print(index)
            res.pop(index)
            # print(res)
            n -= 1
        return res[0]
