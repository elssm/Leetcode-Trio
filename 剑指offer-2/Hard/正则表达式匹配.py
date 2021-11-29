class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        if len(s) == 0:
            if len(p) == 0 or len(p) == 2 * p.count('*'):
                return True
            else:
                return False
        if len(s) != 0 and len(p) == 0: return False
        # 下面只剩下s和p均不为空串的情况了
        return ((s[-1] == p[-1] or p[-1] == '.') and self.isMatch(s[:-1], p[:-1])) \
               or (p[-1] == '*' and (
                    ((s[-1] == p[-2] or p[-2] == '.') and self.isMatch(s[:-1], p)) or self.isMatch(s, p[:-2])))
