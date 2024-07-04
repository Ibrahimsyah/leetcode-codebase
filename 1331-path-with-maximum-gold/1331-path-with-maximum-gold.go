func getMaximumGold(grid [][]int) int {
    visited := make([][]bool, len(grid))
    for i := range visited {
        visited[i] = make([]bool, len(grid[i]))
    }

    max := 0
    var f func(i, j int) int
    f = func(i, j int) int {
        if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || visited[i][j] || grid[i][j] == 0 {
            return 0
        }

        visited [i][j] = true
        up := f(i-1, j)
        down := f(i+1, j)
        left := f(i, j-1)
        right := f(i, j+1)

        max := up
        if down > max {
            max = down
        }
        if left > max {
            max = left
        }
        if right > max {
            max = right
        }

        visited [i][j] = false
        return grid[i][j] + max
    }

    for i := range grid {
        for j := range grid[0] {
            res := f(i, j)
            if res > max {
                max = res
            }
        }
    }

    return max
}
