func reverse(x int) int {
    s := 1
    var r int64
    if x < 0 {
        s = -1
        x *= s
    }

    for {
        if x <= 0 {
            break
        }

        r = r*10 + int64(x%10)
        x /= 10
    }

    if r < math.MinInt32 || r > math.MaxInt32 {
        return 0
    }

    return int(r) * s 
}