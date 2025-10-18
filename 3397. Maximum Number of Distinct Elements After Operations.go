var INF = int(1e9)

func maxDistinctElements(nums []int, k int) int {
    slices.Sort(nums)

    v := [][]int{}
    for _, num := range nums {
        v = append(v, []int{num - k, num + k})
    }

    last := -INF
    result := 0
    for i := 0; i < len(v); i++ {
        next := max(last + 1, v[i][0])

        if next <= v[i][1] {
            last = next
            result++
        }
    }
    return result
}
