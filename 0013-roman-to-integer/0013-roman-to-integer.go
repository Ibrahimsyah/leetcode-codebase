var romanMap = map[string]int {
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
}

func romanToInt(s string) int {
    result := 0
    index := 0
    for index < len(s){
        i := index
        index++
        char := string(s[i])
        if i == len(s) - 1 {
            result += romanMap[char] 
            break
        }
        switch char {
            case "I":
                if string(s[i+1]) == "V" {
                    result += 4
                    index++
                    continue
                }
                if string(s[i+1]) == "X" {
                    result += 9
                    index++
                    continue
                }
            case "X":
                if string(s[i+1]) == "L" {
                    result += 40
                    index++
                    continue
                }
                if string(s[i+1]) == "C" {
                    result += 90
                    index++
                    continue
                }
            case "C":
                if string(s[i+1]) == "D" {
                    result += 400
                    index++
                    continue
                }
                if string(s[i+1]) == "M" {
                    result += 900
                    index++
                    continue
                }
        }
        result += romanMap[char] 
    }

    return result
}