class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        # 初始化列表
        res = ['']
        for i in range(n):
            tmp = set()
            # 遍历res中的结果，对每一个结果的每一位插入'()'，利用set去重复
            for s in res:
                for j in range(len(s) + 1):
                    tmp.add(s[:j] + '()' + s[j:])
            res = list(tmp)
        return res
