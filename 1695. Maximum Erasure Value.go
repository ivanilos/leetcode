func maximumUniqueSubarray(nums []int) int {
    freq := map[int]int{}

    left := 0
    curSum := 0
    result := 0
    for right := 0; right < len(nums); right++ {
        curSum += nums[right]
        freq[nums[right]]++

        for left <= right && freq[nums[right]] > 1 {
            freq[nums[left]]--
            curSum -= nums[left]
            left++
        }

        result = max(result, curSum)
    }

    return result
}
