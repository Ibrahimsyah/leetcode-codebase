func tribonacci(n int) int {
    if n == 0 {
        return 0
    }

    if n < 3 {
        return 1
    }

    t1, t2, t3 := 0, 1, 1
    for i := 3; i <= n; i++ {
        temp := t1 + t2 + t3
        t1 = t2
        t2 = t3
        t3 = temp
    }
    
    return t3
}