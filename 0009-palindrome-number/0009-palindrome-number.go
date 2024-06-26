func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    r := 0
    c := x
    for c != 0 {
        r = r * 10 + c % 10
        c /= 10
    }

    return r == x
}