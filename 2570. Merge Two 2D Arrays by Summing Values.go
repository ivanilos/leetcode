func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    result := [][]int{}

    pos1 := 0
    pos2 := 0

    for pos1 < len(nums1) && pos2 < len(nums2) {
        if nums1[pos1][0] == nums2[pos2][0] {
            sum := nums1[pos1][1] + nums2[pos2][1]
            result = append(result, []int{nums1[pos1][0], sum})
            pos1++
            pos2++
        } else if nums1[pos1][0] < nums2[pos2][0] {
            result = append(result, []int{nums1[pos1][0], nums1[pos1][1]})
            pos1++
        } else {
            result = append(result, []int{nums2[pos2][0], nums2[pos2][1]})
            pos2++
        }
    }

    for pos1 < len(nums1) {
        result = append(result, []int{nums1[pos1][0], nums1[pos1][1]})
        pos1++
    }
    for pos2 < len(nums2) {
        result = append(result, []int{nums2[pos2][0], nums2[pos2][1]})
        pos2++
    }
    
    return result
}
