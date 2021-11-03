class Solution(object):
    def maxSlidingWindow(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        if len(nums) == 0:
            return []
        res = []
        for i in range(len(nums) - k + 1):
            res.append(max(nums[i:k + i]))
        return res
