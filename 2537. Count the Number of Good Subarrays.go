func countGood(nums []int, k int) int64 {
    cur := int64(0)

    result := int64(0)

    left := 0
    freq := map[int]int{}
    for right := 0; right < len(nums); right++ {
        cur += int64(freq[nums[right]])
        freq[nums[right]]++


        for cur >= int64(k) && left < right {
            result += int64(len(nums) - right)

            cur -= int64(freq[nums[left]] - 1)
            freq[nums[left]]--
            left++
        }
    }

    return result
}
