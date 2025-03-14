func maximumCandies(candies []int, k int64) int {
    low := 1
    high := int(1e8)
    result := 0

    for low <= high {
        mid := (low + high) / 2

        if canShare(candies, mid, k) {
            result = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return result
}

func canShare(candies []int, singleUnit int, k int64) bool {
    got := int64(0)

    for _, candy := range candies {
        got += int64(candy / singleUnit)
    }

    return got >= k
}
