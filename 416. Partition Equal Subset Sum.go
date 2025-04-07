func canPartition(nums []int) bool {
    sum := getSum(nums)

    if sum % 2 != 0 {
        return false
    }

    dp := make([][]bool, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = make([]bool, sum + 1)
    }

    for i := 0; i < len(nums); i++ {
        dp[i][0] = true
    }
    dp[0][nums[0]] = true


    result := false
    for i := 1; i < len(nums); i++ {
        for j := 1; j <= sum; j++ {
            dp[i][j] = dp[i - 1][j]
            if j - nums[i] >= 0 {
                dp[i][j] = dp[i][j] || dp[i - 1][j - nums[i]]
            }

            if j == sum / 2 && dp[i][j] == true {
                return true
            }
        }
    }

    return result
}

func getSum(nums []int) int {
    result := 0

    for _, num := range nums {
        result += num
    }
    return result
}
