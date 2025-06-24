func findKDistantIndices(nums []int, key int, k int) []int {
    n := len(nums)
    kDistant := map[int]bool{}

    for i := 0; i < n; i++ {
        if nums[i] == key {
            for j := max(0, i - k); j <= min(n - 1, i + k); j++ {
                kDistant[j] = true
            }
        }
    }

    result := []int{}
    for key, _ := range kDistant {
        result = append(result, key)
    }

    slices.Sort(result)

    return result
}
