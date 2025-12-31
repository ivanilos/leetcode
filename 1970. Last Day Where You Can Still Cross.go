func getIdx(row, col, rows, cols int) int {
    return  (row - 1) * cols + (col - 1)
}

func getReverseIdx(idx, rows, cols int) []int {
    return []int{(idx / cols) + 1,  (idx % cols) + 1}
}

func canReach(rows int, cols int, floodedIds int, cells [][]int) bool {
    floodedCellsMap := map[int]bool{}

    for i := 0; i <= floodedIds; i++ {
        cell := cells[i]
        idx := getIdx(cell[0], cell[1], rows, cols)
        floodedCellsMap[idx] = true
    }

    visited := map[int]bool{}
    for j := 1; j <= cols; j++ {
        curCellIdx := getIdx(1, j, rows, cols)
        if _, ok := floodedCellsMap[curCellIdx]; ok {
            continue
        }
        DFS(1, j, rows, cols, floodedCellsMap, visited)
    }

    for j := 1; j <= cols; j++ {
        curCellIdx := getIdx(rows, j, rows, cols)
        if visited[curCellIdx] {
            return true
        }
    }

    return false
}

func DFS(row, col, rows, cols int, floodedCellsMap map[int]bool, visited map[int]bool) {
    idx := getIdx(row, col, rows, cols)

    visited[idx] = true

    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if i != 0 && j != 0 {
                continue
            }

            newRow := row + i
            newCol := col + j
            newIdx := getIdx(newRow, newCol, rows, cols)

            if 1 <= newRow && newRow <= rows && 1 <= newCol && newCol <= cols {
                _, flooded := floodedCellsMap[newIdx]
                _, hasVisited := visited[newIdx]

                if !flooded && !hasVisited {
                    DFS(newRow, newCol, rows, cols, floodedCellsMap, visited)
                }
            }
        }
    }
}

func latestDayToCross(row int, col int, cells [][]int) int {
    low := 0
    high := len(cells) - 1
    best := -1

    for low <= high {
        mid := (low + high) / 2

        if canReach(row, col, mid, cells) {
            low = mid + 1
            best = mid
        } else {
            high = mid - 1
        }
    }

    return best + 1
}
