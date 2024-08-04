func rangeSum(nums []int, n int, left int, right int) int {
    sums := []int{}

    for start := 0; start < len(nums); start++ {
        sum := 0
        for end := start; end < len(nums); end++ {
            sum += nums[end]
            sums = append(sums, sum)
        }
    }

    sort.Ints(sums)

    result := 0
    const MOD = int(1e9) + 7
    for i := left - 1; i <= right - 1; i++ {
        result = (result + sums[i]) % MOD
    }
    return result
}
