func countBadPairs(nums []int) int64 {
    freq := map[int]int{}

    // j - i != nums[j] - nums[i] <=> nums[i] - i != nums[j] - j
    result := int64(0)
    for j, num := range nums {
        result += int64(j - freq[num - j])

        freq[num - j]++
    }
    return result
}
