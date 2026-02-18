func hasAlternatingBits(n int) bool {
    if n <= 2 {
        return true
    }
    return (n ^ (n >> 1)) == nextPowerOf2(n) - 1
}

func nextPowerOf2(n int) int {
    n--
    n = n | (n >> 1)
    n = n | (n >> 2)
    n = n | (n >> 4)
    n = n | (n >> 8)
    n = n | (n >> 16)
    n++

    return n
}
