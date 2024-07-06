func majorityElement(nums []int) int {
    current, count := nums[0], 1
    for _, num := range nums {
        if num != current {
            count --
        } else {
            count ++ 
        }

        if count == 0 {
            current = num
            count = 1
        }
    }

    return current
}