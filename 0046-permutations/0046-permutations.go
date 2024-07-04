func permute(nums []int) [][]int {
    res, temp := [][]int{}, []int{}
    used := make([]bool, len(nums))

    var b func()
    b = func (){
        if len(temp) == len(nums) {
            c := make([]int, len(temp))
            copy(c, temp)
            res = append(res, c)
            return
        }

        for i, num := range nums {
            if used[i] {
                continue
            }
            used[i] = true
            temp = append(temp, num)
            b()
            used[i] = false
            temp = temp[:len(temp)-1]
        }
    }

    b()
    return res
}