const INF = int(1e9)

func splitArray(nums []int, k int) int {
    low := 0
    high := INF
    result := low

    for low <= high {
        mid := (low + high) / 2

        if canSplit(nums, k, mid) {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return result
}

func canSplit(nums []int, k int, maxSum int) bool {
    curSum := 0
    splits := 0
    for _, num := range nums {
        if num > maxSum {
            return false
        }

        if curSum + num > maxSum {
            splits++
            curSum = num
        } else {
            curSum += num
        }
    }
    splits++

    return splits <= k
}
