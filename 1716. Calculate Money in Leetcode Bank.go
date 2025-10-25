func totalMoney(n int) int {
    weeks := n / 7
    leftoverDays := n % 7

    // first week = ((1 + 7) * 7) / 2 = 28
    // each week = 28 / 35 / 42 / ...
    // with W weeks = (28 + (28 + (W - 1) * 7)) * (W) / 2
    result := ((28 + (28 + (weeks - 1) * 7)) * weeks) / 2

    // leftoverDays
    for i := 0; i < leftoverDays; i++ {
        result += weeks + 1 + i
    }

    return result
}
