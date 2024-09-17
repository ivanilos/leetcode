func uncommonFromSentences(s1 string, s2 string) []string {
    words1 := strings.Fields(s1)
    words2 := strings.Fields(s2)

    freq := map[string]int{}
    for _, word := range words1 {
        freq[word]++
    }
    for _, word := range words2 {
        freq[word]++
    }

    result := []string{}
    for key, val := range freq {
        if val == 1 {
            result = append(result, key)
        }
    }
    return result
}
