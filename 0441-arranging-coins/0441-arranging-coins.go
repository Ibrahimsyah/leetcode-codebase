func arrangeCoins(n int) int {
    if n == 0 {
        return 0
    }

    res := 1
    for n > 0 {
        if n - res < 0 {
            return res - 1
        }
        n -= res
        res++
    }

    return res - 1
}