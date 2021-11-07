class Solution(object):
    def validateStackSequences(self, pushed, popped):
        """
        :type pushed: List[int]
        :type popped: List[int]
        :rtype: bool
        """
        s = []  # 初始化一个栈
        for i in popped:
            # 得到每一个pop出去的数字在push数组中的下标
            # 加一为了方便后续对s栈进行插入操作，即这里res-1才是真正的下标
            res = pushed.index(i) + 1
            # 如果当前pop出去的数在push数组的下标大于s栈的长度
            # 则需要将下标数字之前的数字都写入s栈中
            if res > len(s):
                for k in range(len(s), res):
                    s.append(k)
                # 这一步对pop出去的下标置为-1
                s[res - 1] = -1
            # 如果当前pop出去的数在push数组的下标小于s栈的长度
            # 则我们只需要判断在s栈中对于res下标之后的数值是否为-1即可
            elif res < len(s):
                for j in range(res, len(s)):
                    # 如果在res下标之后存在不为-1的下标，则直接返回false
                    if s[j] != -1:
                        return False
                        # 如果res下标之后存在s栈中的都是-1
                # 则对当前下标进行置为-1操作
                s[res - 1] = -1
        # 循环结束的话返回True
        return True
