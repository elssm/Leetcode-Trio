class Solution(object):
    def maxCoins(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        # 前后填1
        nums = [1] + nums + [1]
        n = len(nums)
        # 准备二维dp数组
        dp = [[0] * n for _ in range(n)]
        # i从第二个数开始遍历
        for i in range(1, n):
            for j in range(i - 1, -1, -1):
                res = 0
                for k in range(j + 1, i):
                    res = max(res, dp[i][k] + dp[k][j] + nums[i] * nums[k] * nums[j])
                dp[i][j] = res
        return dp[-1][0]
