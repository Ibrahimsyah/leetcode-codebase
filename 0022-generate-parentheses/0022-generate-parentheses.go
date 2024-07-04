func generateParenthesis(n int) []string {
    res, part := []string{}, []string{}
    var b func(i, j int)
    b = func (i, j int) {
        if len(part) == 2*n {
            res = append(res, strings.Join(part, ""))
            return
        }

        if i < n {
            part = append(part, "(")
            b(i+1, j)
            part = part[:len(part)-1]
        }

        if i > j {
            part = append(part, ")")
            b(i, j + 1)
            part = part[:len(part)-1]
        }
    }

    b(0, 0)
    return res
}