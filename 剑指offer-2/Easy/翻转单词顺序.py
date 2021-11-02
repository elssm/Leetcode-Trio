class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        res = ""
        s = s.split(' ')
        s.reverse()
        for i in s:
            if len(i) != 0:
                res += i
                res += " "

        return res[:-1]
