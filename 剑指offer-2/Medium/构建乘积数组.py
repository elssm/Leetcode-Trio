class Solution(object):
    def constructArr(self, a):
        """
        :type a: List[int]
        :rtype: List[int]
        """
        """
        从前向后累乘一次并保存结果，从后向前累乘一次并保存结果
        就[1,2,3,4,5]而言
        """
        if len(a) == 0:
            return []
        res1 = []
        res2 = []
        result = []
        m = 1
        n = 1
        for i in a[:-1]:  # 从前向后乘只计算到倒数第二个数
            m *= i
            res1.append(m)  # res1中保存[1,1*2,1*2*3,1*2*3*4]
        for i in a[::-1][:-1]:  # 从后向前乘只计算到第二个数
            n *= i
            res2.append(n)  # res2中保存[5,5*4,5*4*3,5*4*3*2]
        res2 = res2[::-1]  # 反转后得[5*4*3*2,5*4*3,5*4,5]
        """
        此时res1 = [1,1*2,1*2*3,1*2*3*4]
            res2 = [5*4*3*2,5*4*3,5*4,5]
        我们可以得到当缺失第一个数时，值为res2[0]
        当缺失最后一个数时，值为res1[-1]
        从缺失第二个数到倒数第二个数
        我们可以得到值为res[i]*res[i-1]
        """
        for i in range(len(res2) + 1):
            if i == 0:
                result.append(res2[0])
            elif i == len(res2):
                result.append(res1[-1])
            else:
                result.append(res2[i] * res1[i - 1])
        return result
