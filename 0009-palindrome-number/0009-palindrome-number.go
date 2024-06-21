func isPalindrome(x int) bool {
    target := x
    if target < 0 || (x != 0 && target % 10 == 0){
        return false
    }

    reversed := 0
    for target != 0 {
        reversed = reversed * 10 + target % 10
        target /= 10
    }

    return reversed == x
}