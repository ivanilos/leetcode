func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
    topX, bottomX := x, x + k - 1
    
    for topX < bottomX {        
        for j := y; j < y + k; j++ {
            grid[topX][j], grid[bottomX][j] = grid[bottomX][j], grid[topX][j]
        }
        topX++
        bottomX--
    }

    return grid
}
