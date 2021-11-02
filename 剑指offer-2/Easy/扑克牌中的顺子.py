class Solution(object):
    def isStraight(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        count = 0
        nums.sort()
        # 统计大小王的个数
        for i in nums:
            if i == 0:
                count += 1
        nums = nums[count:]
        # 如果有重复数字直接返回False
        if len(nums) != len(list(set(nums))):
            return False
        for i in range(1, len(nums)):
            count -= (nums[i] - nums[i - 1] - 1)
            if count < 0:
                return False
        return True
