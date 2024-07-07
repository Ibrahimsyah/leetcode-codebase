func addDigits(num int) int {
    for num >= 10 {
        counter := 0
        n := num
        for n > 0 {
            counter += n%10
            n /= 10
        }
        num = counter
    }

    return num
}