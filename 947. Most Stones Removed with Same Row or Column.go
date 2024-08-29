type Pair struct {
    row, col int
}

func removeStones(stones [][]int) int {
    usedRows := map[int]bool{}
    usedCols := map[int]bool{}

    rows := make([][]int, 10005)
    cols := make([][]int, 10005)

    for _, stone := range stones {
        rows[stone[0]] = append(rows[stone[0]], stone[1])
        cols[stone[1]] = append(cols[stone[1]], stone[0])
    }


    result := 0
    for i := 0; i < len(rows); i++ {
        if _, ok := usedRows[i]; ok {
            continue
        }
        changed := true

        rowQueue := []int{i}
        usedRows[i] = true
        colQueue := []int{}

        total := map[Pair]bool{}
        for changed == true {
            changed = false

            for len(rowQueue) > 0 {
                nextRow := rowQueue[0]
                rowQueue = rowQueue[1 : len(rowQueue)]

                for j := 0; j < len(rows[nextRow]); j++ {
                    pos := Pair{nextRow, rows[nextRow][j]}
                    total[pos] = true

                    if _, ok := usedCols[rows[nextRow][j]]; !ok {
                        colQueue = append(colQueue, rows[nextRow][j])
                        usedCols[rows[nextRow][j]] = true
                        changed = true
                    }
                }
            }

            for len(colQueue) > 0 {
                nextCol := colQueue[0]
                colQueue = colQueue[1 : len(colQueue)]

                for j := 0; j < len(cols[nextCol]); j++ {
                    pos := Pair{cols[nextCol][j], nextCol}
                    total[pos] = true

                    if _, ok := usedRows[cols[nextCol][j]]; !ok {
                        rowQueue = append(rowQueue, cols[nextCol][j])
                        usedRows[cols[nextCol][j]] = true
                        changed = true
                    }
                }
            }
        }
        result += max(0, len(total) - 1)
        total = map[Pair]bool{}
    }
    return result
}
