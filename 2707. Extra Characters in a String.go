func minExtraChar(s string, dictionary []string) int {
    wordMap := map[string]bool{}

    for _, word := range dictionary {
        wordMap[word] = true
    }

    result := solve(s, wordMap)
    return result
}

func solve(s string, wordMap map[string]bool) int {
    dp := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = -1
    }

    result := calc(0, wordMap, s, dp)
    return result
}

func calc(pos int, wordMap map[string]bool, s string, dp []int) int {
    if pos >= len(s) {
        return 0
    }
    if dp[pos] != -1 {
        return dp[pos]
    }

    result := len(s)
    for end := pos; end < len(s); end++ {
        result = min(result, end - pos + 1 + calc(end + 1, wordMap, s, dp))

        if _, ok := wordMap[s[pos : end + 1]]; ok {
            result = min(result, calc(end + 1, wordMap, s, dp))
        }
    }
    dp[pos] = result
    return result
}
