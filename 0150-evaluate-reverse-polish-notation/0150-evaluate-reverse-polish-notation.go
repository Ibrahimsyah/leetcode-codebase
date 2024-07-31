func evalRPN(tokens []string) int {
    var stk []int
    for _, n := range tokens {
        if n != "+" && n != "-" && n != "*" && n != "/" {
            val, _ := strconv.Atoi(n)
            stk = append(stk, val)
            continue 
        }
        
        i1 := stk[len(stk) - 2]
        i2 := stk[len(stk) - 1]
        res := 0
        switch n {
            case "+": res = i1 + i2
            case "-": res = i1 - i2
            case "*": res = i1 * i2
            case "/": res = i1 / i2
        }

        stk = stk[:len(stk) - 2]
        stk = append(stk, res)
    }
    
    return stk[len(stk) - 1]
}