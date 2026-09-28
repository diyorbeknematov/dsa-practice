func isValidSudoku(board [][]byte) bool {
    rows := [9][9]bool{}
    cols := [9][9]bool{}
    boxes := [9][9]bool{}

    for i := 0; i < 9; i ++ {
        for j := 0; j < 9; j ++ {
            if board[i][j] == '.' {
                continue
            }
            
            indx := i/3 * 3 + j/3
            digit := board[i][j] - '1'

            if boxes[indx][digit] || 
                rows[i][digit] ||
                cols[j][digit] {
                    return false
                }
            boxes[indx][digit] = true
            rows[i][digit] = true
            cols[j][digit] = true
        }
    }

    return true 
}
