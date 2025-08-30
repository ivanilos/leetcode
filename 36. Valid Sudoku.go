func isValidSudoku(board [][]byte) bool {
    if !validRows(board) {
        return false
    }
    if !validCols(board) {
        return false
    }
    if !validAll3By3(board) {
        return false
    }
    return true
}

func validRows(board [][]byte) bool {
    for row := 0; row < len(board); row++ {
        if !validRow(board, row) {
            return false
        }
    }
    return true
}

func validRow(board [][]byte, row int) bool {
    seen := map[byte]bool{}
    for j := 0; j < len(board[row]); j++ {
        if board[row][j] != '.' {
            if _, ok := seen[board[row][j]]; ok {
                return false
            }
            seen[board[row][j]] = true
        }
    }
    return true
}

func validCols(board [][]byte) bool {
    for col := 0; col < len(board[0]); col++ {
        if !validCol(board, col) {
            return false
        }
    }
    return true
}

func validCol(board [][]byte, col int) bool {
    seen := map[byte]bool{}
    for i := 0; i < len(board); i++ {
        if board[i][col] != '.' {
            if _, ok := seen[board[i][col]]; ok {
                return false
            }
            seen[board[i][col]] = true
        }
    }
    return true
}

func validAll3By3(board [][]byte) bool {
    for i := 1; i < len(board); i += 3 {
        for j := 1; j < len(board[i]); j += 3 {
            if !valid3By3(board, i, j) {
                return false
            }
        }
    }
    return true
}

func valid3By3(board [][]byte, cx, cy int) bool {
    seen := map[byte]bool{}

    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            px, py := cx + i, cy + j
            if board[px][py] != '.' {
                if _, ok := seen[board[px][py]]; ok {
                    return false
                }
                seen[board[px][py]] = true
            }
        }
    }
    return true
}
