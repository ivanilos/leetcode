func maxCoins(nums []int) int {
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)

    dp := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = make([]int, len(nums))
        for j := 0; j < len(nums); j++ {
            dp[i][j] = -1
        }
    }

    return solve(1, len(nums) - 2, nums, dp)
}

func solve(left, right int, nums []int, dp [][]int) int {
    if left > right {
        return 0
    }
    if dp[left][right] != -1 {
        return dp[left][right]
    }

    dp[left][right] = 0
    for i := left; i <= right; i++ {
        score := nums[left - 1] * nums[i] * nums[right + 1]
        dp[left][right] = max(dp[left][right], score + solve(left, i - 1, nums, dp) + solve(i + 1, right, nums, dp))
    }

    return dp[left][right]
}
