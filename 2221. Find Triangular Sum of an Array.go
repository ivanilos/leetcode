func triangularSum(nums []int) int {
    for i := 0; i < len(nums) - 1; i++ {
        for j := 0; j < len(nums) - i - 1; j++ {
            nums[j] = (nums[j] + nums[j + 1]) % 10
        }
    }

    return nums[0]
}
