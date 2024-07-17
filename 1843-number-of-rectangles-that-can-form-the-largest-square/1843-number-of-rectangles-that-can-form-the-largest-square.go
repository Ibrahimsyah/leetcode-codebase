func countGoodRectangles(rectangles [][]int) int {
    max := 0
    for _, rect := range rectangles {
        if min(rect[0], rect[1]) > max {
            max = min(rect[0], rect[1])
        }
    }

    res := 0
    for _, rect := range rectangles {
        if min(rect[0], rect[1]) == max {
            res++
        }
    }

    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}