func merge(nums1 []int, m int, nums2 []int, n int)  {
    nextIdx1 := m - 1
    nextIdx2 := n - 1
    nextPos := m + n - 1

    for nextIdx1 >= 0 && nextIdx2 >= 0 {
        if nums1[nextIdx1] > nums2[nextIdx2] {
            nums1[nextPos] = nums1[nextIdx1]
            nextIdx1--
        } else {
            nums1[nextPos] = nums2[nextIdx2]
            nextIdx2--
        }
        nextPos--
    }
    for nextIdx2 >= 0 {
        nums1[nextPos] = nums2[nextIdx2]
        nextIdx2--
        nextPos--
    }
}
