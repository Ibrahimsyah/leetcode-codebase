func canMakeSquare(grid [][]byte) bool {
    valid := func(i, j int) bool {
        w, b := 0, 0
        if grid[i][j] == 'W' {
            w ++ 
        } else {
            b ++
        }
        if grid[i+1][j] == 'W' {
            w ++ 
        } else {
            b ++
        }
        if grid[i][j+1] == 'W' {
            w ++ 
        } else {
            b ++
        }
        if grid[i+1][j+1] == 'W' {
            w ++ 
        } else {
            b ++
        }

        return w == 3 || w == 4 || b == 3 || b == 4
    }

    return valid(0, 0) || valid(0, 1) || valid(1, 0) || valid(1, 1)
}