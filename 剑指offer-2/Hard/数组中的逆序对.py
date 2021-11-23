import bisect


class Solution(object):
    def reversePairs(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        """
        通过不断地插入值来寻找前面大于当前值的元素个数
        """
        res = 0  # 统计计数
        sorted_nums = []  # 准备一个待插入的列表
        for i in range(len(nums)):
            # 找到nums中当前数在sorted_nums中要插入的下标
            k = bisect.bisect(sorted_nums, nums[i])
            # 当前下标减去要插入的下标就是在大于当前数的个数
            res += i - k
            sorted_nums[k:k] = [nums[i]]  # 插入当前元素
            # bisect.insort(sorted_nums,nums[i]) #插入当前元素
        return res
