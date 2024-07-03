func checkValid(matrix [][]int) bool {
    m := make(map[string]struct{})
    for i := range matrix {
        for j := range matrix[i] {
            if _, f := m[fmt.Sprintf("row-%d-%d", i, matrix[i][j])]; f {
                return false
            }
            if _, f := m[fmt.Sprintf("col-%d-%d", j, matrix[i][j])]; f {
                return false
            }

            m[fmt.Sprintf("row-%d-%d", i, matrix[i][j])] = struct{}{}
            m[fmt.Sprintf("col-%d-%d", j, matrix[i][j])] = struct{}{}
        }
    }
    
    return true
}