func minimumOperations(nums []int) int {
    const MAX_INT = 100
    freq := make([]int, MAX_INT + 1)

    for i := len(nums) - 1; i >= 0; i-- {
        if freq[nums[i]] > 0 {
            pos := i + 1
            return (pos + 2) / 3
        }
        freq[nums[i]]++
    }
    return 0
}
