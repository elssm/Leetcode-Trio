class Solution(object):
    def canJump(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        # 分别计算每一个位置能到达的最远距离
        res = 0
        for i in range(len(nums)):
            # 如果当前下标已经大于之前最远能跳跃的长度，直接返回False
            if i > res:
                return False
            # 如果最远能到的位置已经大于等于数组长度，直接返回True
            if res >= len(nums) - 1:
                return True
            # 更新能够到达的最远距离
            res = max(res, i + nums[i])
        return True


