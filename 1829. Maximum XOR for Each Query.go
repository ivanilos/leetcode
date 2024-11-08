func getMaximumXor(nums []int, maximumBit int) []int {
    curXor := 0

    for _, num := range nums {
        curXor ^= num
    }

    result := []int{}
    for i := len(nums) - 1; i >= 0; i-- {
        result = append(result, solve(curXor, maximumBit))
        curXor ^= nums[i]
    }
    return result
}

func solve(curXor int, maximumBit int) int {
    lastKBits := curXor & ((1 << maximumBit) - 1)
    return ((1 << maximumBit) - 1) ^ lastKBits
}
