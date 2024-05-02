func findMaxK(nums []int) int {
    sort.Ints(nums)

    left := 0
    right := len(nums) - 1
    for left < right {
        if abs(nums[left]) > nums[right] {
            left++
        } else if abs(nums[left]) < nums[right] {
            right--
        } else if nums[left] == - nums[right] {
            return nums[right]
        } else {
            break
        }
    }

    return -1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
