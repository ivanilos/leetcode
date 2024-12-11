func maximumBeauty(nums []int, k int) int {
    sort.Ints(nums)

    l := 0
    result := 0
    for r := 0; r < len(nums); r++ {
        for l < r && nums[r] - k > nums[l] + k {
            l++
        }
        result = max(result, r - l + 1)
    }
    return result
}
