func getLongestSubsequence(words []string, groups []int) []string {
    lastGroup := -1

    result := []string{}
    for i := 0; i < len(words); i++ {
        if groups[i] != lastGroup {
            lastGroup = groups[i]
            result = append(result, words[i])
        }
    }

    return result
}
