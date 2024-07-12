func reverseOnlyLetters(s string) string {
    arr := strings.Split(s, "")
    l, r := 0, len(arr) - 1
    for l < r {
        c1 := []byte(strings.ToLower(arr[l]))[0]
        if c1 < 97 || c1 > 122 {
            l++
            continue
        }
        
        c2 := []byte(strings.ToLower(arr[r]))[0]
        if c2 < 97 || c2 > 122 {
            r--
            continue
        }

        arr[l], arr[r] = arr[r], arr[l]
        l++
        r--
    }

    return strings.Join(arr, "")
}