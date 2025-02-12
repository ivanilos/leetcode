func longestSubarray(nums []int) int {
    result := 0

    left := 0
    zeroes := 0
    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 {
            zeroes++
        }

        for left <= right && zeroes >= 2 {
            if nums[left] == 0 {
                zeroes--
            }
            left++
        }

        result = max(result, right - left)
    }

    return min(len(nums) - 1, result)
}
