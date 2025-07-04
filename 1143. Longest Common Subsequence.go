func longestCommonSubsequence(text1 string, text2 string) int {
    n := len(text1)
    m := len(text2)

    dp := make([][]int, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m + 1) 
    }

    dp[0][0] = 0
    for i := 0; i <= n; i++ {
        dp[i][0] = 0
    }
    for j := 0; j <= m; j++ {
        dp[0][j] = 0
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
            if text1[i - 1] == text2[j - 1] {
                dp[i][j] = max(dp[i][j], 1 + dp[i - 1][j - 1])
            }
        }
    }

    return dp[n][m]
}
