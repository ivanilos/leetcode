func largestRectangleArea(heights []int) int {
    heights = append([]int{0}, heights...)

    st := []int{0}

    result := 0
    for i := 1; i < len(heights); i++ {
        lastHeight := heights[st[len(st) - 1]]

        for lastHeight > heights[i] {
            leftIdx := st[len(st) - 2] + 1
            result = max(result, lastHeight * (i - leftIdx))
            st = st[: len(st) - 1]

            lastHeight = heights[st[len(st) - 1]]
        }
        
        st = append(st, i)
    }

    rightIdx := len(heights)
    for i := 1; i < len(st); i++ {
        leftIdx := st[i - 1] + 1
        result = max(result, (rightIdx - leftIdx) * heights[st[i]])
    }

    return result
}
