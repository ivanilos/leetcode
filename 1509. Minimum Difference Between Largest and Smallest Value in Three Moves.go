func minDifference(nums []int) int {
    MAX_CHANGES := 3
    if len(nums) <= MAX_CHANGES + 1 {
        return 0
    }

    slices.Sort(nums)
    result := nums[len(nums) - 1] - nums[0]
    for i := 0; i <= MAX_CHANGES; i++ {
        result = min(result, nums[(len(nums) - 1) - MAX_CHANGES + i] - nums[i])
    }

    return result
}
