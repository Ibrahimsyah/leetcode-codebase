func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    n1, n2 := 1, 2
    for i := 3; i <= n; i++ {
        temp := n1+n2
        n1 = n2
        n2 = temp
    }

    return n2
}