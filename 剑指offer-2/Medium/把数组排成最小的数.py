class Solution(object):
    def minNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: str
        """
        """
        思路分析
        依次比较nums中的数字，将尽可能组合起来小的数放在res数组的前面
        """
        # 先准备一个空的数组，用来存nums中的每一个数字
        res = []
        flag = 0
        for i in nums:
            # 如果res为空，那么直接添加当前数
            if not res:
                res.append(i)
                continue
            # 如果res不为空，则从nums的当前数与res中的每一个数组合起来进行比较
            for j in range(len(res)):
                flag = 0
                """
                分别将当前数字和res中的每一个数字组合进行比较
                例如nums当前数是30，res中的当前数是3
                我们会得到两个组合，分别是303和330
                将这两个数进行比较，发现303会小一些，所以我们将3放在30后面
                在res数组中表现为继续向后比较
                """
                s1 = str(res[j]) + str(i)
                s2 = str(i) + str(res[j])
                if int(s1) < int(s2):
                    continue  # 继续向后比较
                else:
                    # 如果nums中的数放在res[j]之前会使得值更小一些
                    # 那么我们就把nums中的当前数插入到res[j]之前这个位置
                    res[j:j] = [i]
                    # flag=1说明已经nums中的值已经插入到res中了，不需要再比较了
                    # 直接退出res循环
                    flag = 1
                    break
            if flag:
                continue
            # 如果flag=0说明res已经遍历结束
            # 这个时候只需要将nums当前值插入到res后面即可
            else:
                res.append(i)

        s = ""
        for i in res:
            s += str(i)
        return s
