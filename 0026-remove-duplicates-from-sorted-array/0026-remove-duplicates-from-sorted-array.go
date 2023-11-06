func removeDuplicates(nums []int) int {
    m := make(map[int]struct{})
    ci := 0
    for _, n := range nums {
        if _, f := m[n]; f {
            continue
        }
        m[n] = struct{}{}
        nums[ci] = n
        ci++
    }

    return len(m)
}