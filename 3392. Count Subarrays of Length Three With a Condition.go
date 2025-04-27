func countSubarrays(nums []int) int {
    result := 0

    for i := 2; i < len(nums); i++ {
        if (nums[i] + nums[i - 2]) * 2 == nums[i - 1] {
            result++
        }
    }
    return result
}
