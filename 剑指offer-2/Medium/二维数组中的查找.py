class Solution(object):
    def findNumberIn2DArray(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        for i in matrix:
            l = 0
            r = len(i) - 1
            while l <= r:
                mid = (l + r) / 2
                if i[mid] > target:
                    r = mid - 1
                elif i[mid] < target:
                    l = mid + 1
                elif i[mid] == target:
                    return True
        return False
