func isPalindrome(x int) bool {
    target := x
    if target < 0 {
        return false
    }

    digits := []int{}
    for target != 0 {
        digits = append(digits, target % 10)
        target /= 10
    }

    for i := range digits {
        if digits[i] != digits[len(digits) - i - 1] {
            return false
        }
    }

    return true
}