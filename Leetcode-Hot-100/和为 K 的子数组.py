class Solution(object):
    def subarraySum(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        # 前缀和
        m = {0: 1}
        count = 0
        res = 0
        for i in nums:
            res += i
            count += m.get(res - k, 0)
            m[res] = m.get(res, 0) + 1
        return count
