func maxAscendingSum(nums []int) int {
    result := 0

    curSum := 0
    last := nums[0] - 1
    for i := 0; i < len(nums); i++ {
        if nums[i] > last {
            curSum += nums[i]
        } else {
            curSum = nums[i]
        }

        last = nums[i]
        result = max(result, curSum)
    }
    return result
}
