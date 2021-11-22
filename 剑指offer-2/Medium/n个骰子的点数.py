# import numpy as np
class Solution(object):
    def dicesProbability(self, n):
        """
        :type n: int
        :rtype: List[float]
        """
        # 卷积做法
        # conv1 = [1.0/6 for i in range(6)]
        # if n==1:
        #     return conv1
        # else:
        #     convN = self.dicesProbability(n-1)
        #     return np.convolve(convN,conv1)

        # 动态规划
        dp = [1.0 / 6 for _ in range(6)]  # 骰子数为1时候的值
        for i in range(2, n + 1):
            # tmp长度为n个骰子能取到的值的个数，全部初始化为0
            tmp = [0 for _ in range(i * 5 + 1)]

            for j in range(len(dp)):  # n-1个骰子取到的所有值
                for k in range(6):  # 加上第n个骰子取得的值
                    tmp[j + k] += dp[j] / 6.0  # 之前的值和当前的值叠加
            dp = tmp  # 得到最新的值赋值给dp
        return dp
