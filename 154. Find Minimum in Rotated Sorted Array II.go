func findMin(nums []int) int {
    left := 0
    right := len(nums) - 1
    result := nums[0]

    for left <= right {
        mid := (left + right) / 2

        result = min(result, nums[mid])

        if nums[mid] < nums[right] {
            right = mid - 1
        } else if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right--
        }
    }

    return result
}
