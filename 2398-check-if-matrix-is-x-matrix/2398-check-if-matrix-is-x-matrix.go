func checkXMatrix(grid [][]int) bool {
    for i := range grid {
        for j := range grid[i] {
            if (j == len(grid) - i - 1 || i == j) {
                if grid[i][j] == 0{
                    return false
                }
            } else {
                if grid[i][j] != 0 {
                    return false
                }
            }
        }
    }
    return true
}