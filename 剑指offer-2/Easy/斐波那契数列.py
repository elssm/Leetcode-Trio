class Solution(object):
    def fib(self, n):
        """
        :type n: int
        :rtype: int
        """
        res = [0, 1]
        for i in range(2, n + 1):
            res.append(res[i - 1] + res[i - 2])
        print(res)
        return res[n] % 1000000007
