func longestCommonPrefix(strs []string) string {
    min := len(strs[0])
    for _, sub := range strs {
        if len(sub) < min {
            min = len(sub)
        }
    }

    i := 0
    for i < min {
        for _, str := range strs {
            if strs[0][i] != str[i] {
                return str[:i]
            }
        }
        i++
    }

    return strs[0][:i]
}