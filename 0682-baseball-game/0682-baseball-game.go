func calPoints(operations []string) int {
    var stk []int
    for _, op := range operations {
        if op == "C" {
            stk = stk[:len(stk) - 1]
            continue
        }

        if op == "D" {
            stk = append(stk, 2*stk[len(stk) - 1])
            continue
        }

        if op == "+" {
            stk = append(stk, stk[len(stk) - 1] + stk[len(stk) - 2])
            continue
        }

        val, _ := strconv.Atoi(op)
        stk = append(stk, val)
    }

    total := 0
    for _, s := range stk {
        total += s
    }
    
    return total
}