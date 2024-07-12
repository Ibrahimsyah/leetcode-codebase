func countConsistentStrings(allowed string, words []string) int {
    m := make(map[string]struct{})
    for _, c := range allowed {
        m[string(c)] = struct{}{}
    }

    res := 0
    for _, word := range words {
        valid := false
        for _, c := range word {
            if _, f := m[string(c)]; !f {
                valid = false
                break
            }
            valid = true
        }
        if valid {
            res ++
        }
    }

    return res
}