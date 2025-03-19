func minOperations(nums []int) int {
    result := 0

    for i := 0; i < len(nums) - 2; i++ {
        if nums[i] == 0 {
            result++
            nums[i] = 1
            nums[i + 1] = 1 - nums[i + 1]
            nums[i + 2] = 1 - nums[i + 2]
        }
    }

    for i := len(nums) - 1; i >= len(nums) - 3; i-- {
        if nums[i] == 0 {
            return -1
        }
    }

    return result
}
