func summaryRanges(nums []int) []string {
    res := []string{}
    i := 0
    start := 0
    for i < len(nums) {
        if i == len(nums) - 1 || nums[i + 1] != nums[i] + 1 {
            if i == start {
                res = append(res, strconv.Itoa(nums[i]))
            } else {
                res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
            }
            start = i + 1
        }

        i++
    }

    return res
}