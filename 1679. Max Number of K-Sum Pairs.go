func maxOperations(nums []int, k int) int {
    sort.Ints(nums)

    left := 0
    right := len(nums) - 1
    result := 0
    for left < right {
        if nums[left] + nums[right] < k {
            left++
        } else if nums[left] + nums[right] > k {
            right--
        } else {
            result++
            left++
            right--
        }
    }
    return result
}
