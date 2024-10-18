func countMaxOrSubsets(nums []int) int {
    maxiOr := 0
    maxCount := 0
    for mask := 1; mask < 1 << len(nums); mask++ {
        curOr := getOr(nums, mask)

        if curOr > maxiOr {
            maxiOr = curOr
            maxCount = 1
        } else if curOr == maxiOr {
            maxCount++
        }
    }
    return maxCount
}

func getOr(nums []int, mask int) int {
    curOr := 0
    for i := 0; i < len(nums); i++ {
        bit := (mask >> i) & 1

        if bit == 1 {
            curOr |= nums[i]
        }
    }
    return curOr
}
