func countOdds(low int, high int) int {
    result := (high - low + 1) / 2

    if low % 2 == 1 && high % 2 == 1 {
        result++
    }
    return result
}
