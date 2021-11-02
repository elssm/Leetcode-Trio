class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        res = []
        i = 0
        j = len(nums) - 1
        while i < j:
            if nums[i] + nums[j] == target:
                res.append(nums[i])
                res.append(nums[j])
                return res
            if nums[i] + nums[j] > target:
                j -= 1
            else:
                i += 1
