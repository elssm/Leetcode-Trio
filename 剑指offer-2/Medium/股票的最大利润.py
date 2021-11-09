class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        max_num = 0  # 记录当前股票最大差值
        if len(prices) == 0:
            return 0
        min_num = prices[0]  # 将第一天的股票作为最小的值
        for i in range(1, len(prices)):
            min_num = min(min_num, prices[i - 1])  # 判断当天和当前最小值中的最小值
            if prices[i] - min_num > max_num:  # 如果当天的股票减去之前的股票最小值大于最大差值
                max_num = prices[i] - min_num  # 则替换最大差值
        return max_num
