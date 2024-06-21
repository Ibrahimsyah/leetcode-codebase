func lengthOfLastWord(s string) int {
    str := strings.TrimSpace(s)
    words := strings.Split(str, " ")
    return len(words[len(words) - 1])
}