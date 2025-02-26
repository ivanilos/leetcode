func maxAbsoluteSum(nums []int) int {
    result := abs(nums[0])

    curPrefixSum := nums[0]
    maxPrefix := max(0, nums[0])
    minPrefix := min(0, nums[0])
    for i := 1; i < len(nums); i++ {
        curPrefixSum += nums[i]

        result = max(result, abs(curPrefixSum - minPrefix), abs(curPrefixSum - maxPrefix))

        maxPrefix = max(maxPrefix, curPrefixSum)
        minPrefix = min(minPrefix, curPrefixSum)
    }

    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
