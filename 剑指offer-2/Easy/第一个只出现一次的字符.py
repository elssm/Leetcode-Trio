import collections


class Solution(object):
    def firstUniqChar(self, s):
        """
        :type s: str
        :rtype: str
        """
        d = collections.OrderedDict()
        for i in s:
            if i in d.keys():
                d[i] += 1
            else:
                d[i] = 1
        for i in d.keys():
            if d[i] == 1:
                return i
        return " "
