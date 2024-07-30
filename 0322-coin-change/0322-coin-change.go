func coinChange(coins []int, amount int) int {
    m := make([]int, amount + 1)
    for i := range m {
        m[i] = amount + 1
    }
    m[0] = 0

    for i := 1; i <= amount; i++ {
        for _, c := range coins {
            sub := i - c
            if sub < 0 {
                continue
            }

            m[i] = min(m[i], m[sub] + 1)
        }
    }

    if m[amount] == amount + 1 {
        return -1
    }

    return m[amount]
}

func min (a, b int) int {
    if a < b {
        return a
    }

    return b
}