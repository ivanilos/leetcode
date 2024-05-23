func beautifulSubsets(nums []int, k int) int {
    n := len(
nums)

    result := 0
    for mask := 1; mask < 1 << n; mask++ {
        forbidden := map[int]bool{}

        goodSubset := true
        for i := 0; i < n; i++ {
            bit := (mask >> i) & 1

            if bit == 1 {
                forbidden[nums[i] + k] = true
                forbidden[nums[i] - k] = true

                if _, ok := forbidden[nums[i]]; ok {
                    goodSubset = false
                    break
                }
            }
        }
        if goodSubset {
            result++
        }
    }
    return result
}
