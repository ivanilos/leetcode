func maxCount(banned []int, n int, maxSum int) int {
    sort.Ints(banned)

    result := 0
    curSum := 0
    nextBannedIdx := 0
    for i := 1; i <= n; i++ {
        for nextBannedIdx < len(banned) && i > banned[nextBannedIdx] {
            nextBannedIdx++
        }

        if nextBannedIdx >= len(banned) || i != banned[nextBannedIdx] {
            if curSum + i <= maxSum {
                curSum += i
                result++
            }
        }
    }

    return result
}
