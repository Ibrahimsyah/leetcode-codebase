func findTargetSumWays(nums []int, target int) int {
    m := make(map[string]int)
    var solve func(index, total int) int
    solve = func(index, total int) int {
        if index == len(nums) {
            if total == target {
                return 1
            }
            return 0
        }
        dpKey := key(index, total)
        if v, f := m[dpKey]; f {
            return v
        }

        m[dpKey] = solve(index + 1, total + nums[index]) + solve(index + 1, total - nums[index])
        return m[dpKey]
    }

    return solve(0, 0)
} 

func key(index, total int) string {
    return fmt.Sprintf("%d-%d", index, total)
}