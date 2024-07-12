func reversePrefix(word string, ch byte) string {
    s := []string{}
    for i, c := range word {
        s = append(s, string(c))
        if string(c) == string(ch) {
            str := ""
            for len(s) > 0 {
                str += s[len(s)-1]
                s = s[:len(s)-1]
            }
            return str + word[i+1:]
        }
    }

    return word
}