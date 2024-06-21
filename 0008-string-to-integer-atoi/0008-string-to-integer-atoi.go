
func myAtoi(s string) int {
    input := s
    input = strings.TrimSpace(input)

    var res int64 = 0
    signer := 1
    numberInitialized := false
    signerSet := false
    for _, c := range input {
        if !signerSet && !numberInitialized && c == 45 {
            signer = -1
            signerSet = true
            continue
        }

        if !signerSet && !numberInitialized && c == 43 {
            signer = 1
            signerSet = true
            continue
        }
        
        c = c - '0'
        if c < 0 || c > 9 {
            break
        }

        numberInitialized = true
        if res * 10 + int64(c) > math.MaxInt32 {
            if signer == - 1 {
                return math.MinInt32
            }

            return math.MaxInt32 * signer
        }
        
        res = res * 10 + int64(c)
    }

    res *= int64(signer)
    if res > math.MaxInt32 {
        return math.MaxInt32
    }

    if res < math.MinInt32 {
        return math.MinInt32
    }
    
    return int(res)
}