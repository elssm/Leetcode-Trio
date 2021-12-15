class Solution(object):
    def findUnsortedSubarray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 1:
            return 0
        new_nums = sorted(nums)
        l = r = -1
        i = 0
        j = len(nums) - 1
        while i < len(nums):
            if new_nums[i] != nums[i]:
                l = i
                break
            i += 1
        while j >= 0:
            if new_nums[j] != nums[j]:
                r = j
                break
            j -= 1
        if l == -1 and r == -1:
            return 0
        return r - l + 1
