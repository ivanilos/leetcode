func maximumDifference(nums []int) int {
    result := 0
    mini := nums[0]

    for i := 1; i < len(nums); i++ {
        result = max(result, nums[i] - mini)
        mini = min(mini, nums[i])
    }

    if result > 0 {
        return result
    }

    return -1
}
