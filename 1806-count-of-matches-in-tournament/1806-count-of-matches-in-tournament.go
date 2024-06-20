func numberOfMatches(n int) int {
    teams := n
    var matches int
    for teams > 1 {
        matches += teams / 2
        
        if teams % 2 == 0 {
            teams /= 2
            continue
        }

        teams = teams/2 + 1
    }

    return matches
}