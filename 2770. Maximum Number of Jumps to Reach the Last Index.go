func maximumJumps(nums []int, target int) int {
    result := make([]int, len(nums))
    for i := 1; i < len(nums); i++ {
        result[i] = -1
    }

    for i := 0; i < len(nums); i++ {
        if result[i] == -1 {
            continue
        }
        for j := i + 1; j < len(nums); j++ {
            delta := nums[j] - nums[i]
            if -target <= delta && delta <= target {
                result[j] = max(result[j], 1 + result[i])
            }
        }
    }

    return result[len(nums) - 1]
}
