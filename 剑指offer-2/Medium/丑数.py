class Solution(object):
    def nthUglyNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        res = [1]
        for i in range(1, n):
            res.append(0)
        p1 = p2 = p3 = 0
        for i in range(1, n):
            res[i] = min(min(res[p1] * 2, res[p2] * 3), res[p3] * 5)
            if res[i] == res[p1] * 2:
                p1 += 1
            if res[i] == res[p2] * 3:
                p2 += 1
            if res[i] == res[p3] * 5:
                p3 += 1
        return res[n - 1]
