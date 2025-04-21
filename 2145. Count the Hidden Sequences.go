func numberOfArrays(differences []int, lower int, upper int) int {
    const INF = int(1e9)
    low := 0
    high := 0
    cur := 0

    for _, diff := range differences {
        cur += diff
        low = min(low, cur)
        high = max(high, cur)
    }

    deltaLower := lower - low
    deltaUpper := upper - (high + deltaLower) + 1

    return max(0, deltaUpper)
}
