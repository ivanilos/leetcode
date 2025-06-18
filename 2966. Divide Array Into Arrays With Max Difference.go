func divideArray(nums []int, k int) [][]int {
    slices.Sort(nums)

    result := [][]int{}
    cur := []int{}
    for _, num := range nums {
        cur = append(cur, num)

        if len(cur) == 3 {
            result = append(result, cur)
            cur = []int{}
        }
    }

    if isGood(result, k) {
        return result
    } else {
        return [][]int{}
    }
}

func isGood(result [][]int, k int) bool {
    for _, arr := range result {
        if arr[len(arr) - 1] - arr[0] > k {
            return false
        }
    }

    return true
}
