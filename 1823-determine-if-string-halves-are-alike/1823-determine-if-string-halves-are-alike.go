func halvesAreAlike(s string) bool {
    v := map[string]struct{}{
        "a": {},
        "i": {},
        "u": {},
        "e": {},
        "o": {},
    }

    s = strings.ToLower(s)
    c := 0
    for i, ch := range s {
        if _, f := v[string(ch)]; !f {
            continue
        }        

        if i < len(s) / 2 {
            c++
        } else {
            c--
        }
    }

    return c == 0
}