func maxAdjacentDistance(nums []int) int {
    result := abs(nums[0] - nums[len(nums) - 1])

    for i := 1; i < len(nums); i++ {
        result = max(result, abs(nums[i] - nums[i - 1]))
    }

    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
