func smallestNumber(n int) int {
    result := 1

    for result < n {
        result = result * 2 + 1
    }
    return result
}
