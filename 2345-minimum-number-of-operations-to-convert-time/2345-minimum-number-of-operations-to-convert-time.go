func convertTime(current string, correct string) int {
    curr := timeToMinutes(current)
    corr := timeToMinutes(correct)

    diff := corr - curr
    ans := make([]int, diff + 1)
    for i := range ans {
        ans[i] = diff
    }
    ans[0] = 0

    incr := []int{1, 5, 15, 60}
    for i := 1; i <= diff; i++ {
        for _, inc := range incr {
            if i - inc < 0 {
                continue
            }

            ans[i] = min(ans[i], ans[i - inc] + 1)
        }
    }

    return ans[diff]
}

func timeToMinutes(time string) int {
    s := strings.Split(time, ":")
    h, _ := strconv.ParseInt(s[0], 10, 64)
    m, _ := strconv.ParseInt(s[1], 10, 64)
    return int(h) * 60 + int(m)
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}