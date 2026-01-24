func minPairSum(nums []int) int {
    slices.Sort(nums)

    result := 0
    for i := 0; i < len(nums) / 2; i++ {
        result = max(result, nums[i] + nums[len(nums) - 1 - i])
    }

    return result
}
