func reverse(x int) int {
    target := x

    var nums []int
    for target != 0 {
        nums = append(nums, target % 10)
        target /= 10
    }

    fmt.Println(nums)
    res := 0
    for _, n:= range nums {
        res = res * 10 + n
        if res > math.MaxInt32 || res < math.MinInt32 {
            return 0
        }
    }

    return res
}