func minRemoval(nums []int, k int) int {
    slices.Sort(nums)

    left := 0
    result := len(nums)
    for right := 0; right < len(nums); right++ {
        for left < right && int64(nums[right]) > int64(nums[left]) * int64(k) {
            left++
        }
        result = min(result, len(nums) - (right - left + 1))
    }

    return result
}
