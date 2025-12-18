func maxProfit(prices []int, strategy []int, k int) int64 {
    originalValue := int64(0)

    for i := 0; i < len(prices); i++ {
        originalValue += int64(strategy[i] * prices[i])
    }

    result := originalValue
    windowValue := int64(0)
    for i := 0; i < k; i++ {
        if i < k / 2 {
            windowValue += int64(-strategy[i] * prices[i])
        } else {
            windowValue += int64(-strategy[i] * prices[i] + prices[i])
        }
    }

    result = max(result, originalValue + windowValue)

    for i := 1; i <= len(prices) - k; i++ {
        windowValue += int64(strategy[i - 1] * prices[i - 1])
        windowValue += int64(-prices[i + k / 2 - 1])
        windowValue += int64(-strategy[i + k - 1] * prices[i + k - 1] + prices[i + k - 1])

        result = max(result, originalValue + windowValue)
    }

    return result
}
