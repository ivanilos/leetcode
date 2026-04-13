func getMinDistance(nums []int, target int, start int) int {
    result := len(nums)

    for i := start; i < len(nums); i++ {
        if nums[i] == target {
            result = min(result, i - start)
            break
        }
    }

    for i := start; i >= 0; i-- {
        if nums[i] == target {
            result = min(result, start - i)
            break
        }
    }

    return result
}
