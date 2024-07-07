func isAnagram(s string, t string) bool {
    m := make(map[string]int)
    for _, c := range s {
        m[string(c)]++
    }

    for _, c := range t {
        m[string(c)]--
    }

    for _, v := range m {
        if v != 0 {
            return false
        }
    }

    return true
}