func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))
    l, r := 0, len(nums) - 1
    for i := range nums {
        val := 0
        if abs(nums[l]) >= abs(nums[r]) {
            val = nums[l]
            l++
        } else if abs(nums[l]) < abs(nums[r]) {
            val = nums[r]
            r--
        }

        res[len(nums) - i - 1] = val * val
    }

    return res
}

func abs(a int) int {
    if a < 0 {
        return a * -1
    }

    return a
}