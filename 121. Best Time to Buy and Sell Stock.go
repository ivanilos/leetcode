func maxProfit(prices []int) int {
    result := 0
    mini := int(1e9)

    for i := 0; i < len(prices); i++ {
        result = max(result, prices[i] - mini)
        mini = min(mini, prices[i])
    }
    return result
}
