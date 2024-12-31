func mincostTickets(days []int, costs []int) int {
    dp := make([]int, 365 + 30 + 1)
    need := make([]bool, 365 + 30 + 1)
    daysPass := []int{1, 7, 30}

    for i := 0; i < len(days); i++ {
        need[days[i]] = true
    }

    dp[0] = 0
    for i := 1; i <= 365 + 30; i++ {
        if !need[i] {
            dp[i] = dp[i - 1]
        } else {
            dp[i] = int(1e9)
        }
        for j := 0; j < len(costs); j++ {
            if i >=  daysPass[j] {
                dp[i] = min(dp[i], dp[i - daysPass[j]] + costs[j])
            }
        }
    }

    return dp[365 + 30]
}
