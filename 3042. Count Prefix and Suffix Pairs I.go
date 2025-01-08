func countPrefixSuffixPairs(words []string) int {
    result := 0
    for i := 0; i < len(words); i++ {
        for j := i + 1; j < len(words); j++ {
            if isPrefixAndSuffix(words[i], words[j]) {
                result++
            }
        }
    }
    return result
}

func isPrefixAndSuffix(a, b string) bool {
    if len(a) > len(b) {
        return false
    }

    return a == b[0 : len(a)] && a == b[len(b) - len(a) : len(b)]
}
