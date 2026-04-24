func furthestDistanceFromOrigin(moves string) int {
    dist := 0
    buffer := 0
    for _, move := range moves {
        if move == 'L' {
            dist--
        } else if move == 'R' {
            dist++
        } else {
            buffer++
        }
    }

    return max(-(dist - buffer), dist + buffer)
}
