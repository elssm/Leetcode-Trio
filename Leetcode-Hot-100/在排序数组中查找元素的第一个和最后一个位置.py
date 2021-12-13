class Solution(object):
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        # bisect模块
        # l = bisect.bisect_left(nums, target) #如果存在返回该元素在左边的位置，不存在返回应该插入的位置
        # r = bisect.bisect_right(nums, target) #如果存在返回该元素右边的位置
        # return [l,r-1] if l!=r else [-1,-1]

        # 二分查找
        if not nums:
            return [-1, -1]
        n = len(nums) - 1
        l, r = 0, n
        res = [-1, -1]
        while l <= r:
            mid = (l + r) // 2
            if nums[mid] == target:
                start = mid - 1
                end = mid + 1
                while start >= 0 and nums[start] == target:
                    start -= 1
                while end <= n and nums[end] == target:
                    end += 1
                res[0] = start + 1
                res[1] = end - 1
                return res
            elif nums[mid] > target:
                r = mid - 1
            else:
                l = mid + 1
        return res
