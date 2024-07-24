func sortJumbled(mapping []int, nums []int) []int {
    sort.SliceStable(nums, func(a, b int) bool {
        v1 := getVal(mapping, nums[a])
        v2 := getVal(mapping, nums[b])
        return v1 < v2
    })
    return nums
}

func getVal(mapping []int, val int) int {
    result := 0

    base := 1
    for i := 0;; i++ {
        result += mapping[(val % 10)] * base
        val /= 10
        base *= 10

        if val == 0 {
            break
        }
    }
    return result
}
