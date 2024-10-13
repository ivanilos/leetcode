const INF int = int(1e9)

func maxProfit(prices []int) int {
    dp := make([][]int, len(prices))
    for i := 0; i < len(prices); i++ {
        dp[i] = []int{-INF, -INF}
    }

    return solve(0, 0, prices, dp)
}

func solve(day int, has int, prices []int, dp [][]int) int {
    if day >= len(prices) {
        return 0
    }
    if dp[day][has] != -INF {
        return dp[day][has]
    }

    // do nothing
    result := solve(day + 1, has, prices, dp)

    // buy if does not have
    if has == 0 {
        result = max(result, -prices[day] + solve(day + 1, 1, prices, dp))
    }

    // sell if has
    if has == 1 {
        result = max(result, prices[day] + solve(day + 1, 0, prices, dp))
    }

    dp[day][has] = result
    return result
}
