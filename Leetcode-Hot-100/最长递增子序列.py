from bisect import bisect


class Solution(object):
    def lengthOfLIS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        #动态规划
        # if len(nums) <= 1:
        #     return len(nums)
        # #初始化dp数组，其中dp[i]表示i之前包含i的最长上升子序列的长度
        # dp = [1 for i in range(len(nums))]
        # #用来保存最大值
        # res = 0
        # for i in range(1,len(nums)):
        #     for j in range(i):
        #         #如果当前i的值大于i之前的值
        #         if nums[i] > nums[j]:
        #             #则对当前dp[i]中的值进行更新
        #             dp[i] = max(dp[i],dp[j]+1)
        #     #如果当前dp[i]的值大于res的值，则更新res
        #     if dp[i] > res:
        #         res = dp[i]
        # return res

        dp = []
        for x in nums:
            if x in dp:
                continue
            pos = bisect(dp, x)
            if pos == len(dp):
                dp.append(x)
            else:
                dp[pos] = x
        print(dp)
        return len(dp)
