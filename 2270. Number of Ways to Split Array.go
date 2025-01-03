func waysToSplitArray(nums []int) int {
    sufSum := make([]int, len(nums))

    sufSum[len(nums) - 1] = nums[len(nums) - 1]
    for i := len(nums) - 2; i >= 0; i-- {
        sufSum[i] = nums[i] + sufSum[i + 1]
    }

    result := 0
    curSum := 0
    for i := 0; i < len(nums) - 1; i++ {
        curSum += nums[i];

        if curSum >= sufSum[i + 1] {
            result++
        }
    }
    return result
}
