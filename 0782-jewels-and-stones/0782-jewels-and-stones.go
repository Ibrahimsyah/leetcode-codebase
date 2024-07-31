func numJewelsInStones(jewels string, stones string) int {
    j := make(map[rune]struct{})
    for _, c := range jewels {
        j[c] = struct{}{}
    }

    res := 0
    for _, c := range stones {
        if _, v := j[c]; v {
            res++
        }
    }

    return res
}