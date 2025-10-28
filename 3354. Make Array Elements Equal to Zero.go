func countValidSelections(nums []int) int {
    prefSum := make([]int, len(nums))

    prefSum[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        prefSum[i] = nums[i] + prefSum[i - 1]
    }

    result := 0
    sufSum := 0
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] == 0 {
            delta := abs(prefSum[i] - sufSum)
            if delta == 0 {
                result += 2
            } else if delta == 1 {
                result++
            }
        } else {
            sufSum += nums[i]
        }
    }

    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
