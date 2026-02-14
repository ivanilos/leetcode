func champagneTower(poured int, query_row int, query_glass int) float64 {
    grid := make([][]float64, query_row + 1)
    for i := 0; i <= query_row; i++ {
        grid[i] = make([]float64, query_row + 1)
    }

    grid[0][0] = float64(poured)
    for row := 0; row < query_row; row++ {
        for glass := 0; glass <= row; glass++ {
            overflow := max(0, grid[row][glass] - 1.0)

            grid[row + 1][glass] += overflow / 2
            grid[row + 1][glass + 1] += overflow / 2
        }
    }

    return min(1.0, grid[query_row][query_glass])
}
