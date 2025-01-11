func canConstruct(s string, k int) bool {
    freq := make([]int, 26)

    for _, c := range s {
        idx := int(c - 'a')
        freq[idx]++
    }

    oddFreq := 0
    for i := 0; i < 26; i++ {
        if freq[i] % 2 != 0 {
            oddFreq += 1
        }
    }

    if oddFreq <= k && len(s) >= k {
        return true
    }
    return false
}
