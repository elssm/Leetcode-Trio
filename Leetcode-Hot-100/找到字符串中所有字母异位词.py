class Solution(object):
    def findAnagrams(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: List[int]
        """
        p = list(p)
        s = list(s)
        p.sort()
        l = len(p)
        i = 0
        res = []
        while i + l <= len(s):
            tmp = s[i:i + l]
            tmp.sort()
            if p == tmp:
                res.append(i)
            i += 1
        return res
