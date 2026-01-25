func minimumDifference(nums []int, k int) int {
    slices.Sort(nums)

    result := nums[len(nums) - 1] - nums[0]

    for i := k - 1; i < len(nums); i++ {
        result = min(result, nums[i] - nums[i - (k - 1)])
    }

    return result
}
