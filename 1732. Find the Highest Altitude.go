func largestAltitude(gain []int) int {
    cur := 0
    result := cur

    for _, g := range gain {
        cur += g
        result = max(result, cur)
    }

    return result
}
