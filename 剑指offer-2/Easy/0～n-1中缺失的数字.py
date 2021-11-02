class Solution(object):
    def missingNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        res = []
        for i in range(len(nums) + 1):
            res.append(-1)
        # print(res)
        for i in nums:
            res[i] = 0
        return res.index(-1)
