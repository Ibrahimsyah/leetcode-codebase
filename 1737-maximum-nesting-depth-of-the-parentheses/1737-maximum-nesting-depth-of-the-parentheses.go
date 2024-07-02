func maxDepth(s string) int {
    d, md := 0, 0
    for _, ch := range s {
        if string(ch) == "(" {
            d ++
            if d > md {
                md = d
            }
        }else if string(ch) == ")" {
            d --
        }
    }
    return md
}