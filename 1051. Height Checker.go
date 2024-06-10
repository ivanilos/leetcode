func heightChecker(heights []int) int {
    expected := append([]int{}, heights...)

    slices.Sort(expected)

    result := 0
    for i := 0; i < len(heights); i++ {
        if heights[i] != expected[i] {
            result++
        }
    }
    return result
}
