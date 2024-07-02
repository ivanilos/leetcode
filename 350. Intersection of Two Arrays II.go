func intersect(nums1 []int, nums2 []int) []int {
    const MAX_VAL = 1000 
    freq1 := make([]int, MAX_VAL + 1)
    freq2 := make([]int, MAX_VAL + 1)

    for _, num := range(nums1) {
        freq1[num]++
    }
    for _, num := range(nums2) {
        freq2[num]++
    }

    result := []int{}
    for i := 0; i <= MAX_VAL; i++ {
        qt := min(freq1[i], freq2[i])
        for j := 1; j <= qt; j++ {
            result = append(result, i)
        }
    }
    return result
}
