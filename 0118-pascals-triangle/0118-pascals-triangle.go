func generate(numRows int) [][]int {
    res := [][]int{}
    for i := 0; i < numRows; i ++ {
        r := make([]int, i + 1)
        for j := 0; j <= i; j++{
            if j == 0 || j == i {
                r[j] = 1
            } else {
                r[j] = res[i-1][j] + res[i-1][j-1]
            }
        }
        res = append(res, r)
    }

    return res
}