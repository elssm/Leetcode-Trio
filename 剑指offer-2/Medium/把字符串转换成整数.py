class Solution(object):
    def strToInt(self, str):
        """
        :type str: str
        :rtype: int
        """
        str = list(str.lstrip())  # 去掉开头多余的空格
        num = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
        if len(str) == 0:
            return 0
        if len(str) == 1:
            if str[0] in num:
                return int(str[0])
            else:
                return 0
        res = 0
        i = 0
        while i < len(str):
            if str[i] != "-" and str[i] != "+" and str[i] not in num:
                return 0
            elif str[i] == '-' or str[i] == '+':  # 处理以‘-’和‘+’开头的
                temp = str[i]  # 保存‘+-’符号
                i += 1
                if str[i] not in num:
                    return 0
                else:
                    while i < len(str):
                        if str[i] in num:
                            res *= 10
                            res += int(str[i])
                            i += 1
                        else:  # 如果第i个字符不是数字直接退出
                            break
                    if temp == '-':
                        if res > pow(2, 31):
                            return -pow(2, 31)
                        return -res
                    else:
                        if res > pow(2, 31) - 1:
                            return pow(2, 31) - 1
                        return res
            elif str[i] in num:  # 处理以字母开头的
                while i < len(str):
                    if str[i] in num:
                        res *= 10
                        res += int(str[i])
                        i += 1
                    else:
                        break
                if res > pow(2, 31) - 1:
                    return pow(2, 31) - 1
                else:
                    return res
