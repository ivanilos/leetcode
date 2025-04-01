func mostPoints(questions [][]int) int64 {
    n := len(questions)
    dp := make([]int64, n)

    dp[n - 1] = int64(questions[n - 1][0])
    for i := n - 2; i >= 0; i-- {        
        take := int64(questions[i][0])
        if i + questions[i][1] + 1 < n {
            take += dp[i + questions[i][1] + 1]
        }

        dp[i] = max(dp[i + 1], take)
    }

    return dp[0]
}
