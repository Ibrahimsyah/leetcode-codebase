func canJump(nums []int) bool {
    step := 0
    for _, n := range nums {
        if step < 0 {
            return false
        } else if n > step {
            step = n
        }
        step --
    }

    return true
}