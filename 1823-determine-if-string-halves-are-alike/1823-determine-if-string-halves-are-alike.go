func halvesAreAlike(s string) bool {
    vowels := map[string]struct{}{
        "a": {},
        "i": {},
        "u": {},
        "e": {},
        "o": {},
    }

    a := s[:len(s)/2]
    b := s[len(s)/2:]
    
    var aC, bC uint8
    for _, c := range a {
        if _, f := vowels[strings.ToLower(string(c))]; f {
            aC++
        } 
    }

    for _, c := range b {
        if _, f := vowels[strings.ToLower(string(c))]; f {
            bC++
        } 
    }

    return aC == bC
}