class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 1:
            return nums[0]
        if len(nums) == 2:
            return max(nums[0], nums[1], nums[0] + nums[1])
        opt = [0 for _ in range(len(nums))]
        opt[0] = nums[0]
        # opt[1] = max(nums[0], nums[1], nums[0] + nums[1])
        for i in range(1, len(nums)):
            A = opt[i - 1] + nums[i]  # 选择当前数字
            B = nums[i]  # 不选择当前数字
            opt[i] = max(A, B)
        return max(opt)
