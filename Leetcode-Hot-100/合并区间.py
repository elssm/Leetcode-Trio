class Solution(object):
    def merge(self, intervals):
        """
        :type intervals: List[List[int]]
        :rtype: List[List[int]]
        """
        intervals = sorted(intervals)
        if len(intervals) <= 1:
            return intervals
        res = []
        tmp = []
        i = 0
        while i < len(intervals):
            if not tmp:
                tmp = intervals[i]
                i += 1
            else:
                if intervals[i][0] >= tmp[0] and intervals[i][0] <= tmp[1]:
                    if intervals[i][1] > tmp[1]:
                        tmp[1] = intervals[i][1]
                    i += 1
                else:
                    res.append(tmp)
                    tmp = []
        if tmp:
            res.append(tmp)
        return res