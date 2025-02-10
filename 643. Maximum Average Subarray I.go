func findMaxAverage(nums []int, k int) float64 {
    curSum := 0

    for i := 0; i < k; i++ {
        curSum += nums[i]
    }

    result := float64(curSum) / float64(k)
    for i := k; i < len(nums); i++ {
        curSum += nums[i]
        curSum -= nums[i - k]

        result = max(result, float64(curSum) / float64(k))
    }

    return result
}
