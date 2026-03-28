var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

func exist(board [][]byte, word string) bool {
    rows := len(board)
    cols := len(board[0])

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if board[i][j] == word[0] {
                board[i][j] = '.'
                if search(i, j, 1, board, word) {
                    return true
                }
                board[i][j] = word[0]
            }
        }
    }

    return false
}

func search(x, y, idx int, board [][]byte, word string) bool {
    if idx >= len(word) {
        return true
    }

    found := false
    for i := 0; i < 4; i++ {
        newX := x + dx[i]
        newY := y + dy[i]

        if !found && isIn(newX, newY, board) && board[newX][newY] == word[idx] {
            board[newX][newY] = '.'
            found = found || search(newX, newY, idx + 1, board, word)
            board[newX][newY] = word[idx]
        }
    }

    return found
}

func isIn(x, y int, board [][]byte) bool {
    rows := len(board)
    cols := len(board[0])

    return 0 <= x && x < rows && 0 <= y && y < cols
}
