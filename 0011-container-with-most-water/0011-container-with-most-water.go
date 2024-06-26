func maxArea(height []int) int {
    l, r := 0, len(height) - 1
    maxArea := 0
    for l < r {
        maxArea = max(maxArea, min(height[l], height[r]) * (r - l))
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }

    return maxArea
}

func max(i, j int) int {
    if i < j {
        return j
    }
    return i
}

func min(i, j int) int{
    if i < j {
        return i
    }

    return j
}