func applyOperations(nums []int) []int {
    for i := 0; i + 1 < len(nums); i++ {
        if nums[i] == nums[i + 1] {
            nums[i] = 2 * nums[i]
            nums[i + 1] = 0
        }
    }

    nextFreePos := 0
    for i := 0; i < len(nums); i++ {
        for nextFreePos < i && nums[nextFreePos] != 0 {
            nextFreePos++
        }

        if nums[i] != 0 && nextFreePos < i {
            nums[nextFreePos] = nums[i]
            nextFreePos++
            nums[i] = 0
        }
    }

    return nums
}
