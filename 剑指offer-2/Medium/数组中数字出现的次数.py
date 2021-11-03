class Solution(object):
    def singleNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        # d = {}
        # res = []
        # for i in range(len(nums)):
        #     if nums[i] in d.keys():
        #         d[nums[i]]+=1
        #     else:
        #         d[nums[i]]=1
        # for k,v in d.items():
        #     if v==1:
        #         res.append(k)
        # return res

        res = []
        # d = dict(Counter(nums))
        for k, v in dict(Counter(nums)).items():
            if v == 1:
                res.append(k)
        return res
