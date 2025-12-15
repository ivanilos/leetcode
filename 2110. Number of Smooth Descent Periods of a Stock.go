func getDescentPeriods(prices []int) int64 {
    last := prices[0] - 2
    curDescent := 0

    result := int64(0)
    for _, price := range prices {
        if price == last - 1 {
            curDescent++
        } else {
            curDescent = 1
        }

        result += int64(curDescent)

        last = price
    }

    return result
}
