const INF = int(1e9)

func findMaxForm(strs []string, m int, n int) int {
    sizes := make([][]int, len(strs))
    for i := 0; i < len(strs); i++ {
        sizes[i] = []int{0, 0}
    }

    for i, s := range strs {
        sizes[i][0], sizes[i][1] = getSizes(s)
    }

    dp := make([][][]int, len(strs) + 1)
    for i := 0; i < len(strs) + 1; i++ {
        dp[i] = make([][]int, m + 1)
        for j := 0; j < m + 1; j++ {
            dp[i][j] = make([]int, n + 1)

            for k := 0; k < n + 1; k++ {
                dp[i][j][k] = 0
            }
        }
    }

    result := 0
    for i := 1; i <= len(strs); i++ {
        zeroes := sizes[i - 1][0]
        ones := sizes[i - 1][1] 

        for j := 0; j <= m; j++ {
            for k := 0; k <= n; k++ {
                dp[i][j][k] = dp[i - 1][j][k]
                if j - zeroes >= 0 && k - ones >= 0 {
                    dp[i][j][k] = max(dp[i][j][k], 1 + dp[i - 1][j - zeroes][k - ones])
                }

                result = max(result, dp[i][j][k])
            }
        }
    }

    return result
}

func getSizes(s string) (int, int) {
    zeroes := 0
    
    for _, c := range s {
        if c == '0' {
            zeroes++
        }
    }

    return zeroes, len(s) - zeroes
}
