func isPalindrome(s string) bool {
    r := regexp.MustCompile("[^a-zA-Z0-9]+")
    ss := r.ReplaceAllString(strings.ToLower(s), "")
    for i := 0; i < len(ss)/2; i++ {
        if ss[i] != ss[len(ss)-i-1] {
            return false
        }
    }
    return true
}