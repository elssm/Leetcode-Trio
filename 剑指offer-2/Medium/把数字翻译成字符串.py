class Solution(object):
    def translateNum(self, num):
        """
        :type num: int
        :rtype: int
        """
        if num < 10:
            return 1
        elif num % 100 < 10 or num % 100 > 25:
            return self.translateNum(num / 10)
        else:
            return (self.translateNum(num / 10) + self.translateNum(num / 100))
