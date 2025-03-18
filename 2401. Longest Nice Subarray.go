func longestNiceSubarray(nums []int) int {
    result := 1

    for start := 0; start < len(nums); start++ {
        curOr := 0

        for i := start; i < len(nums); i++ {
            newCurOr := curOr | nums[i]

            if newCurOr == (curOr ^ nums[i]) {
                curOr = newCurOr
                result = max(result, i - start + 1)
            } else {
                break
            }
        }
    }

    return result
}
