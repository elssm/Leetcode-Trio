class Solution(object):
    def printNumbers(self, n):
        """
        :type n: int
        :rtype: List[int]
        """
        res = []
        for i in range(1, 10 ** n):
            res.append(i)
        return res
