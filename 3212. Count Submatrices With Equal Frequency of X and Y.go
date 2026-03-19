func numberOfSubmatrices(grid [][]byte) int {
    xFreq := make([][]int, len(grid))
    yFreq := make([][]int, len(grid))

    for i := 0; i < len(grid); i++ {
        xFreq[i] = make([]int, len(grid[0]))
        yFreq[i] = make([]int, len(grid[0]))
    }

    result := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 'X' {
                xFreq[i][j]++
            } else if grid[i][j] == 'Y' {
                yFreq[i][j]++
            }

            if i > 0 {
                xFreq[i][j] += xFreq[i - 1][j]
                yFreq[i][j] += yFreq[i - 1][j]
            }
            if j > 0 {
                xFreq[i][j] += xFreq[i][j - 1]
                yFreq[i][j] += yFreq[i][j - 1]
            }
            if i > 0 && j > 0 {
                xFreq[i][j] -= xFreq[i - 1][j - 1]
                yFreq[i][j] -= yFreq[i - 1][j - 1]
            }

            if xFreq[i][j] > 0 && xFreq[i][j] == yFreq[i][j] {
                result++
            }
        }
    }

    return result
}
