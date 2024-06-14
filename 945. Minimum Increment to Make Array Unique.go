func minIncrementForUnique(nums []int) int {
    slices.Sort(nums)

    result := 0
    prev := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] <= prev {
            result += prev - nums[i] + 1
            prev++
        } else {
            prev = nums[i]
        }
    }
    return result
}
