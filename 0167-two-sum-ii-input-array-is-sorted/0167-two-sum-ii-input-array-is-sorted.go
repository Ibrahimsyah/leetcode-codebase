func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers) - 1
    for l < r {
        res := numbers[l] + numbers[r]
        if res == target {
            return []int{l + 1, r + 1}
        }

        if res < target {
            l++
        } else {
            r--
        }
    }

    return []int{}
}