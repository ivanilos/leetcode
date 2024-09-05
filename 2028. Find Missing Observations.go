func missingRolls(rolls []int, mean int, n int) []int {
    m := len(rolls)
    total := m + n
    sumRolls := getSum(rolls)

    needSum := mean * total - sumRolls

    minSum := n
    maxSum := 6 * n

    if minSum <= needSum && needSum <= maxSum {
        result := make([]int, n)
        for i := 0; i < n; i++ {
            result[i] = 1
        }

        curSum := minSum
        for i := 0; i < n; i++ {
            delta := min(5, needSum - curSum)
            curSum += delta
            result[i] += delta
        }
        return result
    } else {
        return []int{}
    }
}

func getSum(rolls []int) int {
    result := 0
    for _, roll := range rolls {
        result += roll
    }
    return result
}
