func lengthOfLongestSubstring(s string) int {
    var (
        r = 1
        ws = 2
    )

    if len(s) == 0 || len(s) == 1 {
        return len(s)
    }

    var i = 0
    for {
        if i + ws > len(s) {
            break
        }

        sstr:
        for {
            if i + ws > len(s) {
                break
            }

            if !isValid(s[i:i+ws]) {
                i++
                break sstr
            }
            
            r = ws
            ws++
        }
    }  

    return r
}

func isValid (s string) bool {
    m := make(map[rune]struct{})
    for _, r := range s {
        if _, f := m[r]; f {
            return false
        }

        m[r] = struct{}{}
    }

    return true
}