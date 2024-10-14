func hIndex(citations []int) int {
    const MAX_CITATIONS int = 1005
    freq := make([]int, MAX_CITATIONS)

    for _, citation := range citations {
        freq[citation]++
    }

    cur := 0
    for i := MAX_CITATIONS - 1; i >= 0; i-- {
        if freq[i] > 0 {
            cur += freq[i]
        }
        if cur >= i {
            return i
        }
    }
    return 0
}
