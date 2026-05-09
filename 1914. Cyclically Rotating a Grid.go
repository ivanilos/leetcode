func rotateGrid(grid [][]int, k int) [][]int {
    rows := len(grid)
    cols := len(grid[0])

    for it := 0; it < rows; it++ {
        if it >= rows - it - 1 || it >= cols - it - 1 {
            break
        }
        rotate(grid, k, it, rows - it - 1, it, cols - it - 1)
    }

    return grid
}

func rotate(grid [][]int, times, minRow, maxRow, minCol, maxCol int) {
    pos := [][]int{}
    values := []int{}

    for i := minRow; i <= maxRow; i++ {
        pos = append(pos, []int{i, minCol})
        values = append(values, grid[i][minCol])
    }
    for j := minCol + 1; j <= maxCol; j++ {
        pos = append(pos, []int{maxRow, j})
        values = append(values, grid[maxRow][j])
    }

    for i := maxRow - 1; i >= minRow; i-- {
        pos = append(pos, []int{i, maxCol})
        values = append(values, grid[i][maxCol])
    }

    for j := maxCol - 1; j > minCol; j-- {
        pos = append(pos, []int{minRow, j})
        values = append(values, grid[minRow][j])
    }
    
    times = times % len(values)

    for i := 0; i < len(values); i++ {
        idx := (i + times) % len(values)
        grid[pos[idx][0]][pos[idx][1]] = values[i]
    }
}
