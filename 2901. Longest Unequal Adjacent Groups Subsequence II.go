func getWordsInLongestSubsequence(words []string, groups []int) []string {
    canAdd := make([][]bool, len(words))
    for i := 0; i < len(words); i++ {
        canAdd[i] = make([]bool, len(words))
    }

    for i := 0; i < len(words); i++ {
        for j := i + 1; j < len(words); j++ {
            canAdd[i][j] = isValidTransition(words[i], groups[i], words[j], groups[j])
        }
    }

    dp := make([]int, len(words))
    chain := make([]int, len(words))
    for i := 0; i < len(words); i++ {
        chain[i] = -1
    }

    result := 0
    resultStart := 0
    for i := len(words) - 1; i >= 0; i-- {
        dp[i] = 1
        for j := i + 1; j < len(words); j++ {
            if canAdd[i][j] && 1 + dp[j] > dp[i] {
                chain[i] = j
                dp[i] = 1 + dp[j]
            }
        }

        if dp[i] > result {
            result = dp[i]
            resultStart = i
        }
    }

    return getSubsequence(resultStart, words, chain)
}

func isValidTransition(wordA string, groupA int, wordB string, groupB int) bool {
    if groupA == groupB {
        return false
    }
    if len(wordA) != len(wordB) {
        return false
    }

    return hamming(wordA, wordB) == 1
}

func hamming(wordA, wordB string) int {
    result := 0

    for i := 0; i < len(wordA); i++ {
        if wordA[i] != wordB[i] {
            result++
        }
    }
    return result
}

func getSubsequence(start int, words []string, chain []int) []string {
    result := []string{words[start]}

    for chain[start] != -1 {
        start = chain[start]
        result = append(result, words[start])
    }

    return result
}
