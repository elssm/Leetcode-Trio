class Solution(object):
    def findTargetSumWays(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """

        # 动态规划
        # if sum(nums) < target or (sum(nums) + target) % 2 == 1: return 0
        # P = (sum(nums) + target) // 2
        # dp = [1] + [0 for _ in range(P)]
        # for num in nums:
        #     for j in range(P,num-1,-1):dp[j] += dp[j - num]
        # return dp[P]

        # 回溯
        def backtrace(nums, target, index, sums):
            nonlocal count
            if index == len(nums):
                if sums == target:
                    count += 1
            else:
                backtrace(nums, target, index + 1, sums + nums[index])
                backtrace(nums, target, index + 1, sums - nums[index])

        count = 0
        backtrace(nums, target, 0, 0)
        return count


