func maxNumberOfBalloons(text string) int {
    m := map[string]int{
        "b": 0,
        "a": 0,
        "l": 0,
        "o": 0,
        "n": 0,
    }

    for _, c := range text {
        if _, f := m[string(c)]; f {
            m[string(c)]++
        }
    }

    total := math.MaxInt32
    for k, v := range m {
        t := v
        if k == "l" || k == "o" {
            t = v / 2
        }

        if t < total {
            total = t
        }
    }

    return total
}