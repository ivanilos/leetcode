func smallestDistancePair(nums []int, k int) int {
    sort.Ints(nums)

    low := 0
    high := nums[len(nums) - 1] - nums[0]
    result := 0

    for low <= high {
        mid := (low + high) / 2

        if lesserDistances(nums, mid) < k {
            result = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return result
}

func lesserDistances(nums []int, dist int) int {
    count := 0
    left := 0
    for right := 0; right < len(nums); right++ {
        for left < right && nums[right] - nums[left] >= dist {
            left++
        }
        count += right - left
    }
    return count
}
