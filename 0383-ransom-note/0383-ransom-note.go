func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[rune]int)
    for _, c := range magazine {
        m[c] ++
    }

    for _, c := range ransomNote {
        if v, f := m[c]; !f || v == 0 {
            return false
        }
        m[c]--
    }

    return true
}