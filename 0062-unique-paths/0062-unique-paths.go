func uniquePaths(m int, n int) int {
    dp := make(map[string]int)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                dp[key(i,j)] = 1
                continue
            }

            dp[key(i,j)] = dp[key(i-1,j)] + dp[key(i, j-1)]
        }
    }

    return dp[key(m-1, n-1)]
}

func key(i, j int) string {
    return fmt.Sprintf("%d%d", i, j)
}