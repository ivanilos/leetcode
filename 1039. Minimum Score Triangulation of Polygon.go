const INF = int(1e9)

func minScoreTriangulation(values []int) int {
    dp := make([][]int, len(values))
    for i := 0; i < len(values); i++ {
        dp[i] = make([]int, len(values))
        for j := 0; j < len(values); j++ {
            dp[i][j] = INF
        }
    }

    return solve(values, 0, len(values) - 1, dp)
}

func solve(values []int, left, right int, dp [][]int) int {
    if right - left <= 1 {
        return 0
    }
    if dp[left][right] != INF {
        return dp[left][right]
    }

    for i := left + 1; i < right; i++ {
        score := values[left] * values[i] * values[right]
        rest := solve(values, left, i, dp) + solve(values, i, right, dp)

        dp[left][right] = min(dp[left][right], score + rest)
    }

    return dp[left][right]
}
