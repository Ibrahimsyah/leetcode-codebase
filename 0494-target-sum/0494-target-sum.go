func findTargetSumWays(nums []int, target int) int {
    var solve func(index, total int) int
    solve = func(index, total int) int {
        if index == len(nums) {
            if total == target {
                return 1
            }
            return 0
        }

        return solve(index + 1, total + nums[index]) + solve(index + 1, total - nums[index])
    }

    return solve(0, 0)
} 