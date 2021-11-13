class Solution(object):
    def cuttingRope(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n <= 3:
            return n - 1
        res = 1
        while n > 4:
            res *= 3
            n -= 3
        return (res * n) % 1000000007
