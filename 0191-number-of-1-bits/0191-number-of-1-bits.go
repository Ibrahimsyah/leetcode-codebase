func hammingWeight(n int) int {
    c := 0
    for n > 0 {
        if n%2 == 1 {
            c++
        }
        n /= 2
    }

    return c
}