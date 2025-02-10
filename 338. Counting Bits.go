func countBits(n int) []int {
    result := make([]int, n + 1)

    result[0] = 0

    for i := 1; i <= n; i++ {
        onBit := i & -i

        result[i] = 1 + result[i - onBit]
    }

    return result
}
