func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers) - 1

    for l < r {
        result := numbers[l] + numbers[r]
        if result == target {
            return []int{l + 1, r + 1}
        }

        if result < target {
            l ++
        } else {
            r -- 
        }
    }

    return []int{}
}