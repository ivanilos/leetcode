func intersection(nums1 []int, nums2 []int) []int {
    const maxNum = 1000
    vecFreq := [maxNum + 1]int{}

    for _, num := range(nums1) {
        vecFreq[num] = 1
    }
    for _, num := range(nums2) {
        if vecFreq[num] == 1 {
            vecFreq[num] = 2
        }
    }

    result := []int{}
    for i := 0; i <= maxNum; i++ {
        if vecFreq[i] == 2 {
            result = append(result, i)
        }
    }
    return result
}
