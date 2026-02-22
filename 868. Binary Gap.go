func binaryGap(n int) int {
    result := 0

    last := n
    for i := 0; (n >> i) > 0; i++ {
        if ((n >> i) & 1) == 1 {
            result = max(result, i - last)
            last = i
        }
    }

    return result
}
