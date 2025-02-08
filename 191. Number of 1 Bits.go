func hammingWeight(n int) int {
    result := 0
    for n > 0 {
        result++
        n -= n & -n
    }
    return result
}
