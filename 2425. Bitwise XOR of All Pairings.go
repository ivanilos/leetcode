func xorAllNums(nums1 []int, nums2 []int) int {
    result := 0

    xor1 := 0
    for i := 0; i < len(nums1); i++ {
        xor1 ^= nums1[i]
    }

    xor2 := 0
    for i := 0; i < len(nums2); i++ {
        xor2 ^= nums2[i]
    }

    if len(nums1) % 2 == 1 {
        result ^= xor2
    }

    if len(nums2) % 2 == 1 {
        result ^= xor1
    }

    return result
}
