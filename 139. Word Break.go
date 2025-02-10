func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s) + 1)

    dp[len(s)] = true
    for i := len(s) - 1; i >= 0; i-- {
        for j := 0; j < len(wordDict) && !dp[i]; j++ {
            word := wordDict[j]
            if len(word) + i <= len(s) && s[i : i + len(word)] == word {
                dp[i] = dp[i] || dp[i + len(word)]
            }
        }
    }
    return dp[0]
}
