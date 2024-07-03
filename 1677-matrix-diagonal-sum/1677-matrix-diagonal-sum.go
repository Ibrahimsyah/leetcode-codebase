func diagonalSum(mat [][]int) int {
    sum := 0
    for i := range mat {
        sum += mat[i][i]
        if len(mat) % 2 == 0 || i != len(mat)/2 {
            sum += mat[i][len(mat) - i - 1]
        }
    }   

    return sum
}