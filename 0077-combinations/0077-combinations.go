func combine(n int, k int) [][]int {
    res, temp := [][]int{}, []int{}
    var b func(index int)
    b = func (index int) {
        if len(temp) == k {
            c := make([]int, k)
            copy(c, temp)
            res = append(res, c)
            return
        }

        for i := index; i <= n; i++ {
            temp = append(temp, i)
            b(i + 1)
            temp = temp[:len(temp) - 1]
        }
    }   

    b(1)
    return res
}