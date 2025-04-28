func countSubarrays(nums []int, k int64) int64 {
    n := int64(len(nums))
    result := int64(n * (n + 1) / 2)

    left := int64(0)
    curSum := int64(0)
    for right := int64(0); right < n; right++ {
        curSum += int64(nums[right])

        for left <= right && (right - left + 1) * curSum >= k {
            result -= n - right
            
            curSum -= int64(nums[left])
            left++
        }
    }

    return result
}
