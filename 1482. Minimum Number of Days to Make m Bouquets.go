func minDays(bloomDay []int, m int, k int) int {
    low := 1
    high := int(1e9)

    result := -1
    for low <= high {
        mid := (low + high) / 2

        curCount := 0
        bouquets := 0
        for i := 0; i < len(bloomDay); i++ {
            if bloomDay[i] <= mid {
                curCount++
                if curCount == k {
                    bouquets++
                    curCount = 0
                }
            } else {
                curCount = 0
            }
        }

        if bouquets >= m {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return result
}
