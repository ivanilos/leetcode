func isGood(nums []int) bool {
    freq := make([]int, len(nums))

    for _, num := range nums {
        if num < len(nums) {
            freq[num]++
        }
    }

    for i := 1; i < len(nums) - 1; i++ {
        if freq[i] != 1 {
            return false
        }
    }

    return freq[len(nums) - 1] == 2
}
