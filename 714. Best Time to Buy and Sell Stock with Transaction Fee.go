var INF = int(1e9)

func maxProfit(prices []int, fee int) int {
    n := len(prices)

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, 2)
        for j := 0; j < 2; j++ {
            dp[i][j] = -INF
        }
    }

    dp[0][0] = 0
    dp[0][1] = -prices[0]
    for i := 1; i < n; i++ {
        dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i] - fee)
        dp[i][1] = max(dp[i - 1][0] - prices[i], dp[i - 1][1])
    }

    return dp[n - 1][0]
}
