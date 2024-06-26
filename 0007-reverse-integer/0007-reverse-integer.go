func reverse(x int) int {
    r, c := 0, x
    for c != 0 {
        r = 10 * r + c%10
        if r > math.MaxInt32 || r < math.MinInt32 {
            return 0
        }
        c /= 10
    }

    return r   
}