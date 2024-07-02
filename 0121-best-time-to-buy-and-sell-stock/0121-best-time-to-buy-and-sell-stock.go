func maxProfit(prices []int) int {
    price := math.MaxInt32
    res := 0
    for _, p := range prices {
        if p < price {
            price = p
        }

        profit := p - price
        if profit > res {
            res = profit
        }
    }

    return res
}