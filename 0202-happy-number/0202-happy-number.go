func isHappy(n int) bool {
    m := map[int]struct{}{
        n: {},
    }
    for n > 1 {
        counter := 0
        num := n
        for num > 0 {
            f := num%10
            counter += f * f
            num /= 10 
        }
    
        n = counter
        if _, f := m[n]; f {
            return false
        }

        m[n] = struct{}{}
    }

    return true
}