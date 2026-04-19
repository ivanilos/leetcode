func maxDistance(nums1 []int, nums2 []int) int {
    result := 0

    nums2Idx := 0
    for i := 0; i < len(nums1); i++ {
        for nums2Idx < i && nums2Idx + 1 < len(nums2) {
            nums2Idx++
        }
        for nums2Idx + 1 < len(nums2) && nums1[i] <= nums2[nums2Idx + 1] {
            nums2Idx++
        }

        if i <= nums2Idx && nums1[i] <= nums2[nums2Idx] {
            result = max(result, nums2Idx - i)
        }
    }

    return result
}
