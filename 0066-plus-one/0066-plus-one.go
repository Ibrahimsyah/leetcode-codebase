func plusOne(digits []int) []int {
    left := 1
    for i := len(digits) - 1; i >= 0; i-- {
        digits[i] += left
        left = 0
        if digits[i] == 10 {
            left = 1
            digits[i] = 0 
        }

        if i == 0 && left == 1 {
            digits = append([]int{1}, digits...)
        }
    }   

    return digits
}