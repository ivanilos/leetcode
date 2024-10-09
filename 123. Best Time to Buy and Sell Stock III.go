const INF int = int(1e9)

func maxProfit(prices []int) int {
    k := 2
    dp := initDP(k, prices)

    return solve(0, 0, 0, k, prices, dp)
}

func initDP(k int, prices []int) [][][]int {
    dp := make([][][]int, len(prices))
    for i := 0; i < len(prices); i++ {
        dp[i] = make([][]int, k)
        for j := 0; j < k; j++ {
            dp[i][j] = make([]int, 2)
            for p := 0; p < 2; p++ {
                dp[i][j][p] = -INF
            }
        }
    }

    return dp
}

func solve(pos int, ops int, hasStock int, maxOps int, prices []int, dp [][][]int) int {
    if pos >= len(prices) || ops >= maxOps {
        return 0
    }
    if dp[pos][ops][hasStock] != -INF {
        return dp[pos][ops][hasStock]
    }

    // do nothing
    result := solve(pos + 1, ops, hasStock, maxOps, prices, dp)

    // buy if does not have stock
    if hasStock == 0 {
        result = max(result, -prices[pos] + solve(pos + 1, ops, 1, maxOps, prices, dp))
    }

    // sell if has stock
    if hasStock == 1 && ops < maxOps {
        result = max(result, prices[pos] + solve(pos + 1, ops + 1, 0, maxOps, prices, dp))
    }

    dp[pos][ops][hasStock] = result
    return result
}
