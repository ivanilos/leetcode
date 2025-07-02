func equalPairs(grid [][]int) int {
    n := len(grid)

    result := 0
    for row := 0; row < n; row++ {
        for col := 0; col < n; col++ {
            match := true
            for pos := 0; pos < n; pos++ {
                if grid[row][pos] != grid[pos][col] {
                    match = false
                    break
                }
            }

            if match {
                result++
            }
        }
    }

    return result
}
