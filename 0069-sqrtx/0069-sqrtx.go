func mySqrt(x int) int {
    l, r := 1, x
    for l <= r {
        m := l + (r - l) / 2
        squared := m * m
        if squared < x {
            l = m + 1
        } else if squared > x {
            r = m - 1
        } else {
            return m
        }
    }

    return r
}