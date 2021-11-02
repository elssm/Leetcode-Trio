class Solution(object):
    def getLeastNumbers(self, arr, k):
        """
        :type arr: List[int]
        :type k: int
        :rtype: List[int]
        """
        # first
        arr.sort()
        return arr[:k]

        # second
        # if k==0:
        #     return []
        # res=arr[:k]
        # res.sort()
        # for i in arr[k:]:
        #     if i >res[-1]:
        #         continue
        #     res[-1]=i
        #     res.sort()
        # return res
