func maximumTripletValue(nums []int) int64 {
    n := len(nums)
    pref := make([][]int, n)
    suf := make([][]int, n)

    pref[0] = []int{nums[0], nums[0]}
    for i := 1; i < n; i++ {
        pref[i] = []int{min(pref[i - 1][0], nums[i]), max(pref[i - 1][1], nums[i])}
    }

    suf[n - 1] = []int{nums[n - 1], nums[n - 1]}
    for i := n - 2; i >= 0; i-- {
        suf[i] = []int{min(suf[i + 1][0], nums[i]), max(suf[i + 1][1], nums[i])}
    }

    result := int64(0)
    for i := 1; i < n - 1; i++ {
        result = max(result, int64(pref[i - 1][0] - nums[i]) * int64(suf[i + 1][0]))
        result = max(result, int64(pref[i - 1][1] - nums[i]) * int64(suf[i + 1][1]))
    }

    return result
}
