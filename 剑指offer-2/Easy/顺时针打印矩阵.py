class Solution(object):
    def spiralOrder(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        # 使用zip将二维数组逆时针旋转90°
        res = []
        while len(matrix) != 0:
            for i in matrix[0]:
                res.append(i)
            matrix.pop(0)
            matrix = map(list, zip(*matrix))[::-1]
        return res
