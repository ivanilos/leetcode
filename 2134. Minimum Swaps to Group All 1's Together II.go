func minSwaps(nums []int) int {
    ones := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            ones++
        }
    }

    curOnesInRotation := 0
    for i := 0; i < ones; i++ {
        if nums[i] == 1 {
            curOnesInRotation++
        }
    }

    result := ones - curOnesInRotation
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] == 1 {
            curOnesInRotation--
        }
        if nums[(ones - 1 + i) % len(nums)] == 1 {
            curOnesInRotation++
        }
        result = min(result, ones - curOnesInRotation)
    }
    return result
}
