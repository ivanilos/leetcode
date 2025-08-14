const INF = int(1e9)

func maxSizeSlices(slices []int) int {
    n := len(slices)
    return max(solve(slices[0 : n - 1], n / 3), solve(slices[1 : n], n / 3))
}

func solve(v []int, need int) int {
    n := len(v)

    dp := make([][]int, need + 1)
    for i := 0; i < need + 1; i++ {
        dp[i] = make([]int, n + 1)
    }

    for i := 0; i <= need; i++ {
        dp[i][0] = -INF
    }
    for j := 0; j < n; j++ {
        dp[1][j + 1] = max(v[j], dp[1][j])
    }


    for got := 2; got <= need; got++ {
        for i := 0; i < n; i++ {
            dp[got][i + 1] = dp[got][i]

            if i - 1 >= 0 {
                dp[got][i + 1] = max(dp[got][i + 1], v[i] + dp[got - 1][i - 1])
            }
        }
    }

    return dp[need][n]
}
