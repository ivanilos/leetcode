var MOD int = int(1e9 + 7)

func checkRecord(n int) int {
    dp := make([][][]int, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = make([][]int, 2)
        for j := 0; j <= 1; j++ {
            dp[i][j] = make([]int, 3)
            for k := 0; k <= 2; k++ {
                dp[i][j][k] = -1
            }
        }
    }

    return solve(n, 0, 0, dp)
}

func solve(n, totalAbsent, consecutiveLate int, dp [][][]int) int {
    if n == 0 {
        return 1
    }
    if dp[n][totalAbsent][consecutiveLate] != -1 {
        return dp[n][totalAbsent][consecutiveLate]
    }

    // present
    result := solve(n - 1, totalAbsent, 0, dp)
    // absent
    if totalAbsent == 0 {
        result = (result + solve(n - 1, totalAbsent + 1, 0, dp)) % MOD
    }
    // late
    if consecutiveLate <= 1 {
        result = (result + solve(n - 1, totalAbsent, consecutiveLate + 1, dp)) % MOD
    }
    dp[n][totalAbsent][consecutiveLate] = result

    return result
}
