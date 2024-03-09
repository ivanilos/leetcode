func getCommon(nums1 []int, nums2 []int) int {
    firstIdx := 0
    secondIdx := 0

    for firstIdx < len(nums1) && secondIdx < len(nums2) {
        if nums1[firstIdx] == nums2[secondIdx] {
            return nums1[firstIdx]
        } else if nums1[firstIdx] < nums2[secondIdx] {
            firstIdx++
        } else {
            secondIdx++
        }
    }
    return -1
}
