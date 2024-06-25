var m = map[string]string{
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
}
func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }

    result := []string{""}
    for _, d := range digits {
        temp := []string{}
        for _, c := range m[string(d)] {
            for _, current := range result {
                temp = append(temp, current + string(c))
            }
        }

        result = temp
    }

    return result
}
