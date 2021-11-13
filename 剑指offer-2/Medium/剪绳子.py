# 绳子段切分的越多，乘积越大，且3做因数比2更优。
class Solution(object):
    def cuttingRope(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n <= 3:
            return n - 1
        sum = 1
        while n > 4:
            sum *= 3
            n -= 3
        return sum * n
