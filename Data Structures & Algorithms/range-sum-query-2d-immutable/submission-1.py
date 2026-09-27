class NumMatrix:

    def __init__(self, matrix: List[List[int]]):
        self.matrix = [[0] * (len(matrix[0]) + 1)
               for _ in range(len(matrix) + 1)]

        for i in range(1, len(self.matrix)):
            for j in range(1, len(self.matrix[0])):
                self.matrix[i][j] = (
                    matrix[i-1][j-1] 
                    + self.matrix[i-1][j]
                    + self.matrix[i][j-1]
                    - self.matrix[i-1][j-1]
                )
        

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        return (
            self.matrix[row2+1][col2+1]
            - self.matrix[row1][col2+1]
            - self.matrix[row2+1][col1]
            + self.matrix[row1][col1]
        )


# Your NumMatrix object will be instantiated and called as such:
# obj = NumMatrix(matrix)
# param_1 = obj.sumRegion(row1,col1,row2,col2)

# 3 0 1 4 2   10 
# 5 6 3 2 1   17
# 1 2 0 1 5   9
# 4 1 0 1 7   13
# 1 0 3 0 5   9

# 2 1 4 3 
# 1 1 2 2 
# 1 2 2 4 