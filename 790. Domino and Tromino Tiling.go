var MOD int = int(1e9 + 7)

func numTilings(n int) int {
    dp := make([][]int, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, 4)
    }

    // dp[pos][freeSpace]
    dp[n][0] = 1
    dp[n][1] = 0
    dp[n][2] = 0

    dp[n - 1][0] = 1
    dp[n - 1][1] = 0
    dp[n - 1][2] = 0

    for i := n - 2; i >= 0; i-- {
        // add vertical, 2 horizontal, or a t-shape rotation
        dp[i][0] = add(dp[i + 1][0], add(dp[i + 2][0], add(dp[i + 1][1], dp[i + 1][2])))
        // add t-shape or horizontal
        dp[i][1] = add(dp[i + 1][2], dp[i + 2][0])
        // add t-shape or horizontal
        dp[i][2] = add(dp[i + 1][1], dp[i + 2][0])
    }
    return dp[0][0]
}

func add(a, b int) int {
    return (a + b) % MOD
}
