func new21Game(n int, k int, maxPts int) float64 {
    if n < k {
        return 0
    }
    dp := make([]float64, n + 1)

    dp[0] = 1

    curSum := 0.0
    for i := 1; i <= n; i++ {
        if i - 1 < k {
            curSum += dp[i - 1]
        }

        if i - maxPts - 1 >= 0 && i - maxPts - 1 < k {
            curSum -= dp[i - maxPts - 1]
        }

        dp[i] = curSum / float64(maxPts)
    }

    result := 0.0
    for i := k; i <= n; i++ {
        result += dp[i]
    }

    return result
}
