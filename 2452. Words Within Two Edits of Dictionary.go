const MAX_MISMATCHES_ALLOWED = 2

func twoEditWords(queries []string, dictionary []string) []string {
    result := []string{}
    for _, query := range queries {
        for _, word := range dictionary {
            if match(query, word) {
                result = append(result, query)
                break
            }
        }
    }

    return result
}

func match(word, pattern string) bool {
    mismatches := 0
    for i := 0; i < len(word); i++ {
        if word[i] != pattern[i] {
            mismatches++
        }
    }

    return mismatches <= MAX_MISMATCHES_ALLOWED
}
