class Solution(object):
    def permutation(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        # 先对s的下标进行全排列
        numlist = []
        for i in range(len(s)):
            numlist.append(i)
        res = []

        def trace(path, choice):
            if len(path) == len(numlist):
                temp = ""
                for j in path:
                    temp += s[j]
                if temp not in res:
                    res.append(temp)
            for i in choice:
                if i in path:
                    continue
                path.append(i)
                trace(path, choice)
                path.pop()

        trace([], numlist)
        return res
