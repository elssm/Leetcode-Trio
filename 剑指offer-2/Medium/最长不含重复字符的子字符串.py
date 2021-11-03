class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        if len(s) == 0:
            return 0
        result = -1  # 保存最长不重复字符串的长度
        t = ""  # 维护一个不重复的字符串
        i = 0
        while i < len(s):
            # 如果当前字符不在t中，则当前字符加入t
            if s[i] not in t:
                t += s[i]
                i += 1
            else:
                # 如果当前字符在t中，判断t的长度是否是最长的
                if len(t) > result:
                    result = len(t)
                t = ""
                # 将当前字符之前的字符串反转
                temp = s[:i][::-1]
                # 从当前字符的上一个重复字符位置开始遍历
                i = len(temp) - temp.index(s[i])
        return max(result, len(t))
