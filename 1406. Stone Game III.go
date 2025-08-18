const INF = int(1e9)

func stoneGameIII(stoneValue []int) string {
    dp := make([]int, len(stoneValue) + 1)

    for i := len(stoneValue) - 1; i >= 0; i-- {
        curSum := 0
        dp[i] = -INF
        for j := i; j < i + 3 && j < len(stoneValue); j++ {
            curSum += stoneValue[j]

            dp[i] = max(dp[i], curSum - dp[j + 1])
        }
    }

    if dp[0] > 0 {
        return "Alice"
    } else if dp[0] < 0 {
        return "Bob"
    } else {
        return "Tie"
    }
}
