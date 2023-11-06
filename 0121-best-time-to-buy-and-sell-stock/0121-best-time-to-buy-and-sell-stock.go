func maxProfit(prices []int) int {
    m := 0
    lp := prices[0]
    for _, p := range prices {
        lp = int(math.Min(float64(lp), float64(p)))
        m = int(math.Max(float64(m), float64(p - lp)))
    }

    return m
}