func romanToInt(s string) int {
    m := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    res := 0
    i := 0
    for i < len(s) {
        if i < len(s) - 1 && m[string(s[i])] < m[string(s[i+1])]{
            res += m[string(s[i+1])] - m[string(s[i])]
            i += 2
        } else {
            res += m[string(s[i])]
            i++
        }
    }

    return res
}