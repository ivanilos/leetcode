func maxDistance(colors []int) int {
    result := 0

    colorToMinIdx := map[int]int{}
    for i := 0; i < len(colors); i++ {
        for pastColor, pastIdx := range colorToMinIdx {
            if colors[i] != pastColor {
                result = max(result, i - pastIdx)
            }
        }

        if _, ok := colorToMinIdx[colors[i]]; !ok && len(colorToMinIdx) < 2 {
            colorToMinIdx[colors[i]] = i
        }
    }

    return result
}
