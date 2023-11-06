func isPalindrome(s string) bool {
    var r string
    for _, c := range s {
        if unicode.IsLetter(c) || unicode.IsDigit(c) {
            r += string(c)
        }
    }
    r = strings.ToLower(r)
    for i := 0; i < len(r)/2; i++ {
        if r[i] != r[len(r)-i-1] {
            return false
        }
    }
    return true
}