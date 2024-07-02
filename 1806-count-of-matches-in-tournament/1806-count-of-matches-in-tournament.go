func numberOfMatches(n int) int {
    m := 0
    for n > 1 {
        if n%2 == 0 {
            m += n / 2
            n /= 2
        } else {
            m += (n-1)/2
            n = (n-1)/2 + 1
        }
    }

    return m
}