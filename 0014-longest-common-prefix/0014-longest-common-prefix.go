func longestCommonPrefix(strs []string) string {
    sort.Slice(strs, func(i, j int) bool { return strs[i] < strs[j]})
    minLength := math.MaxInt32
    for _, s := range strs {
        if len(s) < minLength {
            minLength = len(s)
        }
    }

    result := ""
    for i := 0; i < minLength; i++ {
        if strs[0][i] != strs[len(strs)-1][i] {
            return result 
        }
        result += string(strs[0][i])
    }

    return result
}