from heapq import *


class MedianFinder(object):

    # 菜鸡的数组排序做法
    # def __init__(self):
    #     """
    #     initialize your data structure here.
    #     """
    #     self.res=[]

    # def addNum(self, num):
    #     """
    #     :type num: int
    #     :rtype: None
    #     """
    #     self.res.append(num)

    # def findMedian(self):
    #     """
    #     :rtype: float
    #     """
    #     self.res=sorted(self.res)
    #     if len(self.res)%2:
    #         return self.res[len(self.res)/2]
    #     else:
    #         return (self.res[len(self.res)/2-1]+self.res[len(self.res)/2])/2.0

    # 大佬的堆解法
    def __init__(self):
        # heapq默认小顶堆
        self.A = []  # 小顶堆，保存较大的一半
        # 实现大顶堆进堆的时候取相反数
        self.B = []  # 大顶堆，保存较小的一半

    def addNum(self, num):
        # 如果不相等，说明A多了一个
        # 把当前数存到A，调整堆之后，返回A中的最小值给B
        if len(self.A) != len(self.B):
            heappush(self.B, -heappushpop(self.A, num))
        # 如果相等，就把当前数给B，调整堆之后，返回B中的最大值给A
        else:
            heappush(self.A, -heappushpop(self.B, -num))

    def findMedian(self):
        # 如果是奇数，直接返回A中的最小值
        # 如果是偶数，则返回两者堆顶的平均值
        # print(self.A)
        # print(self.B)
        return self.A[0] if len(self.A) != len(self.B) else (self.A[0] - self.B[0]) / 2.0

# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
