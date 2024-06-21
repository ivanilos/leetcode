func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    maxTime := len(grumpy)

    dp := make([][]int, maxTime)
    for i := 0; i < maxTime; i++ {
        dp[i] = make([]int, 2)
        for j := 0; j < 2; j++ {
            dp[i][j] = -1
        }
    }

    return solve(0, 0, customers, dp, grumpy, minutes, maxTime)
}

func solve(timer int, usedSecret int, customers []int, dp [][]int, grumpy []int, minutes int, maxTime int) int {
    if timer >= maxTime {
        return 0
    }
    if dp[timer][usedSecret] != -1 {
        return dp[timer][usedSecret]
    }

    dp[timer][usedSecret] = 0
    if grumpy[timer] == 0 {
        dp[timer][usedSecret] = customers[timer] + solve(timer + 1, usedSecret, customers, dp, grumpy, minutes, maxTime)
    } else {
        dp[timer][usedSecret] = solve(timer + 1, usedSecret, customers, dp, grumpy, minutes, maxTime)

        // use Secret
        if usedSecret == 0 {
            cur := 0
            for i := timer; i < timer + minutes && i < maxTime; i++ {
                cur += customers[i]
            }
            dp[timer][usedSecret] = max(dp[timer][usedSecret], cur + solve(timer + minutes, 1, customers, dp, grumpy, minutes, maxTime))
        }
    }
    return dp[timer][usedSecret]
}
