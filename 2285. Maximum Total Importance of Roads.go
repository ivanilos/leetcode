func maximumImportance(n int, roads [][]int) int64 {
    freq := make([]int64, n)

    for _, road := range(roads) {
        freq[road[0]]++
        freq[road[1]]++
    }

    slices.Sort(freq)

    result := int64(0)
    for i := n - 1; i >= 0; i-- {
        result += int64(i + 1) * freq[i]
    }
    return result
}
