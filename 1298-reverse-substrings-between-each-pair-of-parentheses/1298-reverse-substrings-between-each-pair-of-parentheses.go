func reverseParentheses(s string) string {
    res := []string{}
    for _, ch := range s {
        if string(ch) == ")" {
            temp := []string{}
            for string(res[len(res)-1]) != "(" {
                temp = append(temp, res[len(res)-1])
                res = res[:len(res)-1]
            }
            res = append(res[:len(res)-1], temp...)
        } else {
            res = append(res, string(ch))
        }
    }

    return strings.Join(res, "")
}