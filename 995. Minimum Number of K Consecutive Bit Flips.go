func minKBitFlips(nums []int, k int) int {
    flip := make([]int, len(nums))

    result := 0
    curElemFlips := 0
    for i := 0; i <= len(nums) - k; i++ {
        curElemFlips += flip[i]
        val := (nums[i] + ((curElemFlips) % 2)) % 2

        if val == 0 {
            result++
            curElemFlips++
            if i + k < len(flip) {
                flip[i + k] = -1
            }
        }
    }
    for i := len(nums) - k + 1; i < len(nums); i++ {
        curElemFlips += flip[i]
        val := (nums[i] + (curElemFlips) % 2) % 2

        if val == 0 {
            return -1
        }
    }
    return result
}
