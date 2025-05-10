func minSum(nums1 []int, nums2 []int) int64 {
    s1, free1 := getSum(nums1)
    s2, free2 := getSum(nums2)

    if s1 + free1 == s2 + free2 {
        return s1 + free1
    } else {
        if s1 + free1 > s2 && free2 == 0 {
            return -1
        }
        if s2 + free2 > s1 && free1 == 0 {
            return -1
        }
        return max(s1 + free1, s2 + free2)
    }
}

func getSum(nums []int) (int64, int64) {
    sum := int64(0)
    zeroes := int64(0)

    for _, num := range nums {
        sum += int64(num)
        if num == 0 {
            zeroes++
        }
    }
    return sum, zeroes
}
