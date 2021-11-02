class Solution(object):
    def findContinuousSequence(self, target):
        """
        :type target: int
        :rtype: List[List[int]]
        """
        i = 1
        j = 2
        res = []
        tmp = []
        while j <= target // 2 + 1:
            if (i + j) * (j - i + 1) / 2 == target:
                for k in range(i, j + 1):
                    tmp.append(k)
                res.append(tmp)
                tmp = []
                i += 1
            elif (i + j) * (j - i + 1) / 2 > target:
                i += 1
            else:
                j += 1
        return res