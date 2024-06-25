func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, n := range nums {
        if index, f := m[target-n]; f {
            return []int{index, i}
        }
        m[n] = i
    }

    return []int{}
}