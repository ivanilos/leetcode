func rob(nums []int) int {
    dp := make([]int, len(nums) + 1)

    dp[len(nums)] = 0
    dp[len(nums) - 1] = nums[len(nums) - 1]
    for i := len(nums) - 2; i >= 0; i-- {
        dp[i] = max(dp[i + 1], nums[i] + dp[i + 2])
    }
    return dp[0]
}
