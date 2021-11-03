class Solution(object):
    def add(self, a, b):
        """
        :type a: int
        :type b: int
        :rtype: int
        """
        # while b!=0: #while循环超时
        #     c = a^b
        #     b = (a&b)<<1
        #     a = c
        # return a
        return a if b == 0 else add(a ^ b, (a & b) << 1)
