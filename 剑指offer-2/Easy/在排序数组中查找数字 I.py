class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        # c=0
        # for i in nums:
        #     if i==target:
        #         c+=1
        # return c

        l = 0
        r = len(nums) - 1
        c = 0
        while l < r:
            m = (l + r) / 2
            if nums[m] >= target:
                r = m
            if nums[m] < target:
                l = m + 1
        while l < len(nums) and nums[l] == target:
            c += 1
            l += 1
        return c
