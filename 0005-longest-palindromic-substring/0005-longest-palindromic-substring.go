func longestPalindrome(s string) string {
    if len(s) == 1 {
        return s
    }

    result := ""
    for i := range s {
        l, r := i, i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if len(s[l:r+1]) > len(result) {
                result = s[l:r+1]
            }
            l--
            r++
        }

        l, r = i, i + 1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if len(s[l:r+1]) > len(result) {
                result = s[l:r+1]
            }
            l--
            r++
        }
    }

    return result
}