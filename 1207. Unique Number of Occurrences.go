func uniqueOccurrences(arr []int) bool {
    freq := map[int]int{}

    for _, val := range arr {
        freq[val]++
    }

    existsFreq := map[int]bool{}
    for _, val := range freq {
        if _, ok := existsFreq[val]; ok {
            return false
        }

        existsFreq[val] = true
    }

    return true
}
