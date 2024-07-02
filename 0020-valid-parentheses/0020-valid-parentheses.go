func isValid(s string) bool {
    m := map[string]string {
        ")": "(",
        "}": "{",
        "]": "[",
    }
    stack := []string{}
    for _, char := range s {
        c, f := m[string(char)];
        if !f {
            stack = append(stack, string(char))
            continue
        }

        if len(stack) == 0 {
            return false
        }

        if stack[len(stack)-1] != c {
            return false
        }

        stack = stack[:len(stack)-1]
    }

    return len(stack) == 0
}