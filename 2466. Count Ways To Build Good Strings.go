func countGoodStrings(low int, high int, zero int, one int) int {
    const MOD int = int(1e9 + 7)
    dp := make([]int, high + 1)

    dp[0] = 1
    for i := 1; i <= high; i++ {
        if i >= zero {
            dp[i] = (dp[i] + dp[i - zero]) % MOD
        }
        if i >= one {
            dp[i] = (dp[i] + dp[i - one]) % MOD
        }
    }

    result := 0
    for i := low; i <= high; i++ {
        result = (result + dp[i]) % MOD
    }
    return result
}
