func missingNumber(nums []int) int {
    sum := 0
    expected := len(nums)*(2 + len(nums) - 1)/2
    for _, n := range nums {
        sum += n
    }

    return expected - sum
}