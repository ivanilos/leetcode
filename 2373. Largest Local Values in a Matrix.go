func largestLocal(grid [][]int) [][]int {
    n := len(grid)

    result := make([][]int, n - 2)
    for i := 0; i < n - 2; i++ {
        result[i] = make([]int, n - 2)
    }

    for i := 0; i < n - 2; i++ {
        for j := 0; j < n - 2; j++ {
            maxi := grid[i][j]
            for k := 0; k < 3; k++ {
                for p := 0; p < 3; p++ {
                    maxi = max(maxi, grid[i + k][j + p])
                }
            }
            result[i][j] = maxi
        }
    }

    return result
}
