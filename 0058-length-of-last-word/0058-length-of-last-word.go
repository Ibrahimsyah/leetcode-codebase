func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    r := 0
    for _, c := range s {
        if string(c) == " " {
            r = 0
        } else {
            r++
        }
    }

    return r
}