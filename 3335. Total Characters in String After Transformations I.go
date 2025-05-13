func lengthAfterTransformations(s string, t int) int {
    const MOD = int(1e9 + 7)

    dp := make([][]int, t + 1)
    for i := 0; i <= t; i++ {
        dp[i] = make([]int, 26)
    }

    for i := 0; i < 26; i++ {
        dp[0][i] = 1
    }

    for j := 1; j <= t; j++ {
        for i := 25; i >= 0; i-- {
            if i == 25 {
                dp[j][i] = (dp[j - 1][0] + dp[j - 1][1]) % MOD
            } else {
                dp[j][i] = dp[j - 1][i + 1]
            }
        }
    }

    result := 0
    for _, c := range s {
        idx := int(c - 'a')
        result = (result + dp[t][idx]) % MOD
    }

    return result
}
