func reverseVowels(s string) string {
    v := map[string]struct{}{
        "a": {},
        "i": {},
        "u": {},
        "e": {},
        "o": {},
    }

    ss := strings.Split(s, "")
    l, r := 0, len(s) - 1
    for l < r {
        if _, f := v[strings.ToLower(ss[l])]; !f {
            l++
            continue
        }
        if _, f := v[strings.ToLower(ss[r])]; !f {
            r--
            continue
        }

        temp := ss[l]
        ss[l] = ss[r]
        ss[r] = temp
        l++
        r--
    }

    return strings.Join(ss, "")
}