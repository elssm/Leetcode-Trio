class Solution(object):
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: List[str]
        :rtype: bool
        """
        i=0
        while wordDict:
            wordDict = wordDict[i:]
            tmp = s
            for j in wordDict:
                if j in tmp:
                    tmp = tmp.replace(j,"0")
            tmp = tmp.replace("0","")
            if len(tmp):
                i+=1
            else:
                return True
        return False


s = "ccaccc"
wordDict = ["cc","ac"]
a = Solution()
print(a.wordBreak(s,wordDict))