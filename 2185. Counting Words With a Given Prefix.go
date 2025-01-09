func prefixCount(words []string, pref string) int {
    result := 0

    for i := 0; i < len(words); i++ {
        if len(words[i]) >= len(pref) && pref == words[i][0 : len(pref)] {
            result++
        }
    }
    return result
}
