func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    l, r := 0, len(s) - 1
    for l < r {
        if !isAlphanumeric(s[l]) {
            l++
            continue
        }

        if !isAlphanumeric(s[r]) {
            r--
            continue
        }
        if s[l] != s[r] {
            return false
        }

        l++
        r--
    }
    
    return true
}

func isAlphanumeric(s byte) bool {
    if s >= 97 && s <= 122 {
        return true
    }

    if s >= 48 && s <= 57 {
        return true
    }

    return false
} 