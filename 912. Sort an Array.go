func sortArray(nums []int) []int {
    MAX_VAL := int(5 * 1e4)
    bucket := make([]int, MAX_VAL * 2 + 1)

    for _, num := range(nums) {
        bucket[num + MAX_VAL]++
    }

    nextPos := 0
    for val := -MAX_VAL; val <= MAX_VAL; val++ {
        for bucket[val + MAX_VAL] > 0 {
            nums[nextPos] = val
            nextPos++
            bucket[val + MAX_VAL]--
        }
    }
    return nums
}
