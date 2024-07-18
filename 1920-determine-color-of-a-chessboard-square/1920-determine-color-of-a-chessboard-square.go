func squareIsWhite(coordinates string) bool {
    row := coordinates[0] - 'a' + 1
    col := coordinates[1] - 48
    
    if row % 2 == 0 {
        return col % 2 != 0
    }
    
    return col % 2 == 0
}