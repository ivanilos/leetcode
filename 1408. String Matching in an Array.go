func stringMatching(words []string) []string {
    result := []string{}
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words); j++ {
            if i != j && isSubstring(words[i], words[j]) {
                result = append(result, words[i])
                break
            }
        }
    }

    return result
}

func isSubstring(a, b string) bool {
    if len(a) > len(b) {
        return false
    }

    for start := 0; start <= len(b) - len(a); start++ {
        if a == b[start : start + len(a)] {
            return true
        }
    }
    return false
}
