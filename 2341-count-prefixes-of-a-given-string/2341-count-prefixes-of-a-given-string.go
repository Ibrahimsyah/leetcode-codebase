func countPrefixes(words []string, s string) int {
    result := 0
    for _, word := range words {
        if len(word) > len(s) {
            continue
        }
        
        if word[:len(word)] == s[:len(word)] {
            result++
        }
    }

    return result
}