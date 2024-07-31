func trap(height []int) int {
    l := len(height)
    maxL := make([]int, l)
    maxR := make([]int, l)

    wallL, wallR := 0, 0

    for i := 0; i < l; i ++ {
        j := l - i - 1
        maxL[i] = wallL
        maxR[j] = wallR
        if height[i] > wallL {
            wallL = height[i]
        }

        if height[j] > wallR {
            wallR = height[j]
        }
    }

    result := 0
    for i, h := range height {
        res := maxL[i]
        if maxR[i] < res {
            res = maxR[i]
        }

        res -= h
        if res > 0 {
            result += res
        } 
    }

    return result
}