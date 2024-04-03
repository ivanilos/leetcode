func exist(board [][]byte, word string) bool {
    n := len(board)
    m := len(board[0])

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if board[i][j] == word[0] {
                used := map[Pair]bool{}
                used[Pair{i, j}] = true
                if solve(i, j, 1, used, board, n, m, word) {
                    return true
                }
                used[Pair{i, j}] = false
            }
        }
    }
    return false
}

func solve(x, y int, idx int, used map[Pair]bool, board [][]byte, n, m int, word string) bool {
    if idx >= len(word) {
        return true
    }

    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if i != 0 && j != 0 {
                continue
            }
            nx := x + i
            ny := y + j
            visited, _ := used[Pair{nx, ny}]
            if isIn(nx, ny, n, m) && board[nx][ny] == word[idx] && !visited {
                used[Pair{nx, ny}] = true
                if solve(nx, ny, idx + 1, used, board, n, m, word) {
                    return true
                }
                used[Pair{nx, ny}] = false
            }
        }
    }
    return false
}

func isIn(x, y, n, m int) bool {
    return 0 <= x && x < n && 0 <= y && y < m
}

type Pair struct {
    first, second int
}
