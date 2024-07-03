func isValidSudoku(board [][]byte) bool {
    m := make(map[string]struct{})
    for i := range board {
        for j := range board[i] {
            if string(board[i][j]) == "." || string(board[i][j]) == "," {
                continue
            }

            _, rF := m[fmt.Sprintf("row-%d-%d", i, board[i][j])]
            _, cF := m[fmt.Sprintf("col-%d-%d", j, board[i][j])]
            _, sF := m[fmt.Sprintf("seg-%d-%d-%d", i/3, j/3, board[i][j])]
            if cF || rF || sF {
                return false
            }   

            m[fmt.Sprintf("row-%d-%d", i, board[i][j])] = struct{}{}
            m[fmt.Sprintf("col-%d-%d", j, board[i][j])] = struct{}{}
            m[fmt.Sprintf("seg-%d-%d-%d", i/3, j/3, board[i][j])] = struct{}{}
        }
    }

    return true
}