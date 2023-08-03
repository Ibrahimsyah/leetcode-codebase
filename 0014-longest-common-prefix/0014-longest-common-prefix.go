func longestCommonPrefix(strs []string) string {
    minimumArrayLen := len(strs[0])
    for _, str := range strs {
        if len(str) < minimumArrayLen {
            minimumArrayLen = len(str)
        }
    }

    substr := ""
    for i := 0; i < minimumArrayLen; i++ {
        if isCharCommon(strs[0][i], strs, i) {
            substr += string(strs[0][i])
        } else {
            return substr
        }
    }
    
    return substr
}

func isCharCommon(c byte, strs []string, charIndex int) bool {
    for _, str := range strs {
        if str[charIndex] != c {
            return false
        }
    }

    return true
}