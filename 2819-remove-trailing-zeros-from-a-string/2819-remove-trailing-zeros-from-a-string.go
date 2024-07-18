func removeTrailingZeros(num string) string {
    r := len(num) - 1
    for {
        if string(num[r]) != "0" {
            return num[:r+1]
        }
        r --
    }

    return num
}