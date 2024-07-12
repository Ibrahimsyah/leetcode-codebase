func similarPairs(words []string) int {
    res := 0
    m := make(map[string]int)
    for _, word := range words {
        arr := make([]string, 26)
        for _, c := range word {
            arr[c - 'a'] = "1"
        }

        key := generateMapKey(arr)
        if val, f := m[key]; f {
            res += val
            m[key]++
        } else {
            m[key] = 1
        }
    }

    return res
}

func generateMapKey(chars []string) string {
    return strings.Join(chars, "-")
}