func minSteps(n int) int {
    dp := make([][]int, n + 1) // dp[copy][screen]
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, n + 1)
    }

    dp[0][1] = 0
    dp[0][0] = 0

    result := solve(0, 1, n, dp)
    return result
}

var INF = int(1e9)

func solve(copy, screen, n int, dp [][]int) int {
    if screen > n || copy > n {
        return INF
    }
    if screen == n {
        return 0
    }

    dp[copy][screen] = INF

    if copy > 0 {
        dp[copy][screen] = 1 + solve(copy, screen + copy, n, dp)
    }
    if copy < screen {
        dp[copy][screen] = min(dp[copy][screen], 1 + solve(screen, screen, n, dp))
    }
    return dp[copy][screen]
}
