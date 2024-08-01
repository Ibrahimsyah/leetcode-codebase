func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }

    a, b := nums[0], max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        temp := max(nums[i] + a, b)
        a = b
        b = temp
    }

    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}