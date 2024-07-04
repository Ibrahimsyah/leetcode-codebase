func combinationSum(candidates []int, target int) [][]int {
    res, temp := [][]int{}, []int{}
    var b func(sum, index int)
    b = func(sum, index int) {
        if sum == target {
            combination := make([]int, len(temp))
            copy(combination, temp)
            res = append(res, combination)
            return
        }

        if sum < target {
            for i := index; i < len(candidates); i++ {
                temp = append(temp, candidates[i])
                b(sum + candidates[i], i)
                temp = temp[:len(temp) - 1]
            }
        }
    }

    b(0, 0)
    return res
}