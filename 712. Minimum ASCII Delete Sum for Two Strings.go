const INF = int(1e9)

func minimumDeleteSum(s1 string, s2 string) int {
    n := len(s1)
    m := len(s2)

    dp := make([][]int, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m + 1)
    }

    dp[n][m] = 0
    for i := n - 1; i >= 0; i-- {
        dp[i][m] = dp[i + 1][m] + int(s1[i])
    }
    for j := m - 1; j >= 0; j-- {
        dp[n][j] = dp[n][j + 1] + int(s2[j])
    }

    for i := n - 1; i >= 0; i-- {
        for j := m - 1; j >= 0; j-- {
            dp[i][j] = INF
            if s1[i] == s2[j] {
                dp[i][j] = dp[i + 1][j + 1]
            }

            dp[i][j] = min(dp[i][j], dp[i + 1][j] + int(s1[i]), 
                            dp[i][j + 1] + int(s2[j]))
        }
    }

    return dp[0][0]
}
