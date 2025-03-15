func minCapability(nums []int, k int) int {
    low := 0
    high := int(1e9)

    result := int(1e9)

    for low <= high {
        mid := (low + high) / 2

        if canRob(nums, k, mid) {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return result
}

func canRob(nums []int, k, maxVal int) bool {
    total := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] <= maxVal {
            total++
            i++
        }
    }

    return total >= k
}
