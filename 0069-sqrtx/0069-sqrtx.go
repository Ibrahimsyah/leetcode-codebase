func mySqrt(x int) int {
    l, r := 1, x
    for l <= r {
        m := l + (r - l) / 2
        sqrt := x / m
        if m > sqrt {
            r = m - 1
        } else {
            if m + 1 > x / (m + 1) {
                return m
            }
            l = m + 1
        }
    }

    return 0
}