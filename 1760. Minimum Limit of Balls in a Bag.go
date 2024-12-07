func minimumSize(nums []int, maxOperations int) int {
    low := 1
    high := int(1e9)
    best := 1

    for low <= high {
        mid := (low + high) / 2

        if getOps(nums, mid) <= maxOperations {
            best = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return best
}

func getOps(nums []int, minVal int) int {
    result := 0
    for _, num := range nums {
        result += (num - 1) / minVal
        result = min(result, int(1e9) + 1)
    }
    return result
}
