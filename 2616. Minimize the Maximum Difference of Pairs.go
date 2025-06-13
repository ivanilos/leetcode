func minimizeMax(nums []int, p int) int {
    slices.Sort(nums)

    low := 0
    high := nums[len(nums) - 1] - nums[0]
    best := high

    for low <= high {
        mid := (low + high) / 2

        if canMakePairs(nums, p, mid) {
            best = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return best
}

func canMakePairs(nums []int, p int, maxVal int) bool {
    total := 0

    for i := 1; i < len(nums); i++ {
        if nums[i] - nums[i - 1] <= maxVal {
            total++
            i++
        }
    }

    return total >= p
}
