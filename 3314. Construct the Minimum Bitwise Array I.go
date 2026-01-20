func minBitwiseArray(nums []int) []int {
    result := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        if nums[i] == 2 {
            result[i] = -1
        } else {
            result[i] = nums[i]
            j := 0
            for ((result[i] >> (j + 1)) & 1) == 1 {
                j++
            }
            result[i] ^= 1 << j
        }
    }

    return result
}
