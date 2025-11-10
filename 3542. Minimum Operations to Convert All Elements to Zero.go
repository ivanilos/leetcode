func minOperations(nums []int) int {
    st := []int{}

    result := 0
    for _, num := range nums {
        for len(st) > 0 && st[len(st) - 1] > num {
            st = st[: len(st) - 1]
        }
        if num == 0 {
            continue
        }

        if len(st) == 0 || st[len(st) - 1] < num {
            result++
            st = append(st, num)
        }
    }

    return result
}
