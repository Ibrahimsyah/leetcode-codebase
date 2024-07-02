func maxArea(height []int) int {
    l, r := 0, len(height) - 1
    max := 0
    for l < r {
        h := height[l]
        if height[r] < h {
            h = height[r]
        }

        area := h * (r - l)
        if area > max {
            max = area
        }

        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }

    return max
}