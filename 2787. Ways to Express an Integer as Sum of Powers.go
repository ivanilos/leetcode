const MOD = int(1e9 + 7)

func numberOfWays(n int, x int) int {
    dp := make([][]int, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, n + 1)
        for j := 0; j <= n; j++ {
            dp[i][j] = -1
        }
    }

    result := solve(n, n, x, dp)

    return result
}

func solve(n, nextNum, x int, dp [][]int) int {
    if n == 0 {
        return 1
    }
    if nextNum <= 0 {
        return 0
    }
    if dp[n][nextNum] != -1 {
        return dp[n][nextNum]
    }

    dp[n][nextNum] = 0
    for i := nextNum; i >= 1; i-- {
        toSub := getPower(i, x)

        if n >= toSub {
            dp[n][nextNum] = (dp[n][nextNum] + solve(n - toSub, i - 1, x, dp)) % MOD
        }
    }

    return dp[n][nextNum]
}

func getPower(base, p int) int {
    result := base
    for i := 1; i < p; i++ {
        result *= base
    }

    return result
}
