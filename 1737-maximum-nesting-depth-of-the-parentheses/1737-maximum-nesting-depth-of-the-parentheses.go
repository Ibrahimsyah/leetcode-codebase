func maxDepth(s string) int {
    var depth, maxDepth uint8
    for _, c := range s {
        if string(c) == "(" {
            depth ++
            if depth > maxDepth {
                maxDepth = depth
            }
        } else if string(c) == ")" {
            depth -- 
        }
    }

    return int(maxDepth)
}