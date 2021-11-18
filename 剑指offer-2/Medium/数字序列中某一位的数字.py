class Solution(object):
    def findNthDigit(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n < 10:
            return n
        i = 1
        num = 9
        temp = 0
        """
        如果n大于9的话，判断n应该在哪个区间
        因为1位数一共有9个
        2位数有90*2=180个
        3位数有900*3=2700个
        因此我们可以知道n位数有9*(10**(i-1)*i)个
        """
        while n > num:
            temp = num  # temp保存上一个num的数
            addnum = 9 * (10 ** i) * (i + 1)
            num += addnum
            i += 1
        res = n - temp  # 减去区间之前的那些数
        """
        此时得到的i如果是3，则说明从三位数开始
        可以得到区间在100-1000之间
        从而算得起始的数是10**(3-1)=100
        """
        start_num = 10 ** (i - 1)
        """
        因为在这个区间每一个数都是i位，因此我们得到除数div
        对于三位数来讲，如果除数div=10，则可以得到数字是在100+10-1=109附近
        """
        div = res / i
        mod = res % i
        new_num = start_num + div - 1  # 这个数字就是区间内加上除数之后的数字
        """
        进一步通过余数判断，如果余数为0，则说明当前得到的数字的最后一个数字就是结果
        如果余数不为0，则说明最终结果在下一个数字，从而根据余数得到最终结果
        """
        if mod == 0:
            return int(str(new_num)[-1])
        else:
            return int(str(new_num + 1)[mod - i - 1])

        # 第一种方法，超时了
        # s=0
        # for i in range(n+1):
        #     s+=len(str(i))
        #     if s>=n+1:
        #         return int(str(i)[n-temp])
