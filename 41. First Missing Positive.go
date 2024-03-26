func firstMissingPositive(nums []int) int {
    changeNegativeAndZeroGreaterThanArrayLen(nums)
    result := getSmallestMissingPositive(nums)

    return result
}

func changeNegativeAndZeroGreaterThanArrayLen(nums []int) {
    for i := 0; i < len(nums); i++ {
        if nums[i] <= 0 {
            nums[i] = len(nums) + 1
        }
    }
}

func getSmallestMissingPositive(nums []int) int {
    for i := 0; i < len(nums); i++ {
        idx := abs(nums[i]) - 1
        if idx < len(nums) && nums[idx] > 0 {
            nums[idx] = -nums[idx]
        }
    }

    result := len(nums) + 1
    for i := 0; i < len(nums); i++ {
        if nums[i] >= 0 {
            return i + 1
        }
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
