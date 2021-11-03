class Solution(object):
    def maxValue(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        for i in range(1, len(grid[0])):
            grid[0][i] += grid[0][i - 1]
        for i in range(1, len(grid)):
            grid[i][0] += grid[i - 1][0]
        # print(grid)
        for i in range(1, len(grid)):
            for j in range(1, len(grid[0])):
                grid[i][j] += max(grid[i - 1][j], grid[i][j - 1])
        return grid[-1][-1]
