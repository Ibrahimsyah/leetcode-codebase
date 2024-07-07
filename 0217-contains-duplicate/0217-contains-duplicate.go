func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{})
    for _, n := range nums {
        if _, f := m[n]; f {
            return true
        }
        m[n] = struct{}{}
    }

    return false
}