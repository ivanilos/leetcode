func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    return solve(nums, upper) - solve(nums, lower - 1)
}

func solve(nums []int, upper int) int64 {

    result := int64(0)
    right := len(nums) - 1
    for left := 0; left < right; left++ {
        for left < right && nums[left] + nums[right] > upper {
            right--
        }

        result += int64(max(0, right - left))
    }
    return result
}
