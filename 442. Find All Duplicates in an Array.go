func findDuplicates(nums []int) []int {
    result := []int{}

    for _, num := range nums {
        val := abs(num)

        if nums[val - 1] < 0 {
            result = append(result, val)
        } else {
            nums[val - 1] = -nums[val - 1]
        }
    }
    return result
}

func abs(x int) int{
    if x < 0 {
        return -x
    }
    return x
}
