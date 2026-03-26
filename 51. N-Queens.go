type state struct {
    n int
    usedPieces int
    row int
    usedCol []bool
    usedDiag map[int]bool
    usedSecDiag map[int]bool
    board [][]byte
}

func newState(n int) state {
    board := make([][]byte, n)
    for i := 0; i < n; i++ {
        board[i] = make([]byte, n)
        for j := 0; j < n; j++ {
            board[i][j] = '.'
        }
    }

    return state{
        n,
        0,
        0,
        make([]bool, n),
        map[int]bool{},
        map[int]bool{},
        board,
    }
}

func solveNQueens(n int) [][]string {
    result := [][]string{}
    DFS(newState(n), &result)

    return result
}

func DFS(curState state, result *[][]string) {
    if curState.usedPieces == curState.n {
        addResult(curState.board, result)
        return
    }

    for j := 0; j < curState.n; j++ {
        if !curState.isValidPlacement(curState.row, j) {
            continue
        }

        curState.setQueen(curState.row, j)
        curState.row++
        DFS(curState, result)
        curState.row--
        curState.unSetQueen(curState.row, j)
    }
}

func (s *state) isValidPlacement(row, col int) bool {
    if s.usedCol[col] {
        return false
    }

    diagIdx := row - col
    secDiagIdx := row + col
    if _, ok := s.usedDiag[diagIdx]; ok {
        return false
    }

    if _, ok := s.usedSecDiag[secDiagIdx]; ok {
        return false
    }

    return true
}

func (s *state) setQueen(row, col int) {
    diagIdx := row - col
    secDiagIdx := row + col

    s.board[row][col] = 'Q'
    s.usedPieces++
    s.usedCol[col] = true
    s.usedDiag[diagIdx] = true
    s.usedSecDiag[secDiagIdx] = true
}

func (s *state) unSetQueen(row, col int) {
    diagIdx := row - col
    secDiagIdx := row + col

    s.board[row][col] = '.'
    s.usedPieces--
    s.usedCol[col] = false
    delete(s.usedDiag, diagIdx)
    delete(s.usedSecDiag, secDiagIdx)
}

func addResult(board [][]byte, result *[][]string) {
    rows := make([]string, len(board))

    for i := 0; i < len(board); i++ {
        rows[i] = string(board[i])
    }
    *result = append(*result, rows)
}
