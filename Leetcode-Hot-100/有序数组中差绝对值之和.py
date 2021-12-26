class Solution(object):
    def getSumAbsoluteDifferences(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        """
        解题思路：
        因为是有序排列，因此后一个减去前一个肯定是正数，无论是否加绝对值
        考虑nums = 2 3 5 6这四个数
        res1 = |2-2|+|2-3|+|2-5|+|2-6|
        从第二个绝对值开始，结果都要取反，不考虑相等的值
        将res1改为
        res1 = 2-2+3-2+5-2+6-2 = sum(nums)-2*len(nums)
        
        res2 = |3-2|+|3-3|+|3-5|+|3-6|
        从第三个绝对值开始，结果都要取反，不考虑相等的值
        将res2改为
        res2 = 3-2+3-3+5-3+6-3，其中为了凑到sum(res1)
        我们将 3-2 凑成 (2-3)+(3-2)*2
        则res2 = sum(nums)-3*len(nums) + (3-2)*2
        
        res3 = |5-2|+|5-3|+|5-5|+|5-6|
        从第四个绝对值开始，结果都要取反，不考虑相等的值
        将res3改为
        res3 = 5-2+5-3+5-5+6-5，其中为了凑到sum(res1)
        我们将 5-2+5-3 凑成 (2-5+3-5)+(5-2+5-3)*2
        则res2 = sum(nums)-5*len(nums) + (5-2+5-3)*2
        
        res4 = |6-2|+|6-3|+|6-5|+|6-6|
        将res4改为
        res4 = 6-2+6-3+6-5+6-6，其中为了凑到sum(res1)
        我们将 6-2+6-3+6-5 凑成 (2-6+3-6+5-6)+(6-2+6-3+6-5)*2
        则res2 = sum(nums)-6*len(nums) + (6-2+6-3+6-5)*2
        
        通过上述分析我们最终可以得到第i个数的结果为
        res[i] = (sum(nums)-nums[i]*len(nums)+(nums[i]*i - sum(1...i))*2
        """
        res = []
        n = len(nums)
        sums = sum(nums)
        tmp_sum = 0
        for i in range(len(nums)):
            res.append(sums - nums[i] * n + (nums[i] * i - tmp_sum) * 2)
            tmp_sum += nums[i]
        return res
