func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

    for _, str := range strs {
        arr := make([]int, 26)
        for _, c := range str {
            arr[c - 'a']++
        }
        k := key(arr)
        m[k] = append(m[k], str)
    }

    res := make([][]string, 0, len(m))
    for _, v := range m {
        res = append(res, v)
    }

    return res
}

func key(k []int) string {
    res := ""
    for _, n := range k {
        res += strconv.Itoa(n) + "-"
    }
    return res
}