const INF = int(1e9)

func maximumAmount(coins [][]int) int {
    n := len(coins)
    m := len(coins[0])

    dp := make([][][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][]int, m)
        for j := 0; j < m; j++ {
            dp[i][j] = []int{-INF, -INF, -INF}
        }
    }

    dp[0][0][0] = coins[0][0]
    dp[0][0][1] = 0

    for i := 1; i < n; i++ {
        dp[i][0][0] = coins[i][0] + dp[i - 1][0][0]
        dp[i][0][1] = max(coins[i][0] + dp[i - 1][0][1], dp[i - 1][0][0])
        dp[i][0][2] = max(coins[i][0] + dp[i - 1][0][2], dp[i - 1][0][1]) 
    }

    for j := 1; j < m; j++ {
        dp[0][j][0] = coins[0][j] + dp[0][j - 1][0]
        dp[0][j][1] = max(coins[0][j] + dp[0][j - 1][1], dp[0][j - 1][0])
        dp[0][j][2] = max(coins[0][j] + dp[0][j - 1][2], dp[0][j - 1][1])
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            dp[i][j][0] = coins[i][j] + max(dp[i - 1][j][0], dp[i][j - 1][0])

            dp[i][j][1] = max(dp[i - 1][j][0], dp[i][j - 1][0])
            dp[i][j][1] = max(dp[i][j][1], coins[i][j] + max(dp[i - 1][j][1], dp[i][j - 1][1]))

            dp[i][j][2] = max(dp[i - 1][j][1], dp[i][j - 1][1])
            dp[i][j][2] = max(dp[i][j][2], coins[i][j] + max(dp[i - 1][j][2], dp[i][j - 1][2]))
        }
    }

    return max(dp[n - 1][m - 1][0], dp[n - 1][m - 1][1], dp[n - 1][m - 1][2])
}
