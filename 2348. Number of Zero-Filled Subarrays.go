func zeroFilledSubarray(nums []int) int64 {
    result := int64(0)
    zeroSequenceLen := int64(0)
    for _, num := range nums {
        if num == 0 {
            zeroSequenceLen++
        } else {
            zeroSequenceLen = 0
        }

        result += zeroSequenceLen
    }

    return result
}
