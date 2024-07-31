func longestConsecutive(nums []int) int {
    m := make(map[int]struct{})
    for _, n := range nums {
        m[n] = struct{}{}
    }

    res := 0
    for k := range m {
        if _, f := m[k - 1]; f {
            continue
        }
        
        temp := 0
        c := k
        _, f := m[c]
        for f {
            temp++
            c++
            _, f = m[c]
        }

        if temp > res {
            res = temp
        }
    }

    return res
}