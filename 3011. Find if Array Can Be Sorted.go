func canSortArray(nums []int) bool {
    result := []int{}

    last := nums[0]
    curSortable := []int{}
    for _, val := range nums {
        if canExchange(last, val) {
            curSortable = append(curSortable, val)
        } else {
            sort.Ints(curSortable)
            result = append(result, curSortable...)
            last = val
            curSortable = []int{val}
        }
    }

    sort.Ints(curSortable)
    result = append(result, curSortable...)

    return isSorted(result)
}

func canExchange(a, b int) bool {
    return onBits(a) == onBits(b)
}

func onBits(a int) int {
    result := 0
    for a > 0 {
        result += a & 1
        a /= 2
    }
    return result
}

func isSorted(v []int) bool {
    for i := 1; i < len(v); i++ {
        if v[i] < v[i - 1] {
            return false
        }
    }
    return true
}
