func maximumEnergy(energy []int, k int) int {
    dp := make([]int, len(energy))

    result := energy[len(energy) - 1]
    for i := len(energy) - 1; i >= len(energy) - k && i >= 0; i-- {
        dp[i] = energy[i]
        result = max(result, dp[i])
    }

    for i := len(energy) - k - 1; i >= 0; i-- {
        dp[i] = energy[i] + dp[i + k]
        result = max(result, dp[i])
    }

    return result
}
