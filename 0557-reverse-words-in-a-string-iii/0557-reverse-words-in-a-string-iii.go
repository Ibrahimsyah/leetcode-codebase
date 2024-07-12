func reverseWords(s string) string {
    arr := strings.Split(s, " ")
    for i := range arr {
        arr[i] = reverseWord(arr[i])
    }
    return strings.Join(arr, " ")
}

func reverseWord(word string) string {
    s := ""
    for i := range word {
        s += string(word[len(word)-i-1])
    }
    return s
}