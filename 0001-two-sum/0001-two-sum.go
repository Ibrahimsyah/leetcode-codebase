func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, n := range nums {
        if rI, f := m[target - n]; f {
            return []int{i, rI}
        }

        m[n] = i
    }

    return []int{}
}