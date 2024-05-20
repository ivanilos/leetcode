func subsetXORSum(nums []int) int {
    n := len(nums)

    result := 0
    for mask := 0; mask < (1 << n); mask++ {
        curXor := 0
        for i := 0; i < n; i++ {
            bit := (mask >> i) & 1
            if bit == 1 {
                curXor ^= nums[i]
            }
        }
        result += curXor
    }
    return result
}
