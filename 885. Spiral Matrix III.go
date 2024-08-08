func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
    dirx := []int{0, 1, 0, -1}
    diry := []int{1, 0, -1, 0}

    steps := 1
    dir := 0

    result := [][]int{[]int{rStart, cStart}}
    for len(result) < rows * cols {
        for i := 0; i < steps; i++ {
            rStart += dirx[dir]
            cStart += diry[dir]

            if isIn(rStart, cStart, rows, cols) {
                result = append(result, []int{rStart, cStart})
            }
        }
        dir = (dir + 1) % 4
        if dir == 0 || dir == 2 {
            steps++
        }
    }
    return result
}

func isIn(curRow, curCol, rows, cols int) bool {
    return 0 <= curRow && curRow < rows && 0 <= curCol && curCol < cols
}
