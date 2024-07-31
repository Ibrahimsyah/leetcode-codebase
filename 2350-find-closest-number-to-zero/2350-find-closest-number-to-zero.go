func findClosestNumber(nums []int) int {
    curr := nums[0]
    for _, num := range nums {
        if abs(num) < abs(curr) || (abs(num) == abs(curr) && num > curr) {
            curr = num
        }
    }

    return curr
}

func abs (a int) int{
    if a > 0 {
        return a
    }
    return -a
}