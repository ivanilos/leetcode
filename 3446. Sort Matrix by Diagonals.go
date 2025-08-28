func sortMatrix(grid [][]int) [][]int {
    for i := 0; i < len(grid); i++ {
        sortDiag(grid, i, 0, true)
    }

    for j := 1; j < len(grid[0]); j++ {
        sortDiag(grid, 0, j, false)
    }

    return grid
}

func sortDiag(grid [][]int, x, y int, desc bool) {
    curDiag := []int{}

    for i, j := x, y; i < len(grid) && j < len(grid[0]); i, j = i + 1, j + 1 {
        curDiag = append(curDiag, grid[i][j])
    }

    slices.Sort(curDiag)
    if desc {
        slices.Reverse(curDiag)
    }

    for i, j, k := x, y, 0; i < len(grid) && j < len(grid[0]); i, j, k = i + 1, j + 1, k + 1 {
        grid[i][j] = curDiag[k]
    }

    return
}
