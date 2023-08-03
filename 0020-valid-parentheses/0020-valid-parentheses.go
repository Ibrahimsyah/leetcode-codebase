var a = map[string]string {
    ")": "(",
    "}": "{",
    "]":  "[",
}
func isValid(s string) bool {
    z := make([]string, 0, len(s))
    for _, ch := range s {
        if v, f := a[string(ch)]; f {
            if len(z) == 0 {
                return false
            }

            if z[len(z) - 1] != v {
                return false
            }

            z = z[:len(z) - 1]
            continue
        }

        z = append(z, string(ch))
    }

    return len(z) == 0
}