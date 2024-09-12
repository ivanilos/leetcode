func countConsistentStrings(allowed string, words []string) int {
    allowedMap := map[rune]bool{}
    for _, c := range allowed {
        allowedMap[c] = true
    }

    result := 0
    for _, word := range words {
        good := true
        for _, c := range word {
            if _, ok := allowedMap[c]; !ok {
                good = false
                break
            }
        }
        if good {
            result++
        }
    }
    return result
}
