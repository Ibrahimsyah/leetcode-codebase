func coinChange(coins []int, amount int) int {
    m := make(map[int]int)
    m[0] = 0
    var solve func(target int, coins []int) int
    solve = func(target int, coins []int) int {
        if v, f := m[target];f {
            return v
        }
        
        if target == 0 {
            return 0
        }

        answer := math.MaxInt32
        for _, c := range coins {
            sub := target - c
            if sub < 0 {
                continue
            }

            answer = min(answer, solve(sub, coins) + 1)
        }
        
        m[target] = answer
        return answer
    }

    answer := solve(amount, coins)
    if answer >= 0 && answer != math.MaxInt32 {
        return answer
    }

    return -1
}

func min (a, b int) int {
    if a < b {
        return a
    }

    return b
}