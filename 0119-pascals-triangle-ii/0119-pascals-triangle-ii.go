func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }

    if rowIndex == 1 {
        return []int{1, 1}
    }

    prev := []int{1, 1}
    for i := 2; i <= rowIndex; i++ {
        temp := make([]int, i + 1)
        for j := range temp {
            if j == 0 || j == len(temp) - 1 {
                temp[j] = 1
                continue
            }
            temp[j] = prev[j-1] + prev[j]
        }
        prev = temp
    }

    return prev
}