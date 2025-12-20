func minDeletionSize(strs []string) int {
    rows := len(strs)
    cols := len(strs[0])

    result := 0
    for col := 0; col < cols; col++ {
        prev := strs[0][col]
        for row := 0; row < rows; row++ {
            if strs[row][col] < prev {
                result++
                break
            }
            prev = strs[row][col]
        }
    }

    return result
}
