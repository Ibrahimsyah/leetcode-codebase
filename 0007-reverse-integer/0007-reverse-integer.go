func reverse(x int) int {
    r := 0
    for x != 0 {
        r = r*10 + x%10
        x /= 10
    }

    if r >= math.MaxInt32 || r < math.MinInt32 {
        return 0
    }

    return r
}