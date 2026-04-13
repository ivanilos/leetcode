var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

type region struct {
    positions []pos
    hasPosInBorder bool
}

type pos struct {
    x int
    y int
}

func solve(board [][]byte)  {
    visited := make([][]bool, len(board))
    for i := 0; i < len(board); i++ {
        visited[i] = make([]bool, len(board[i]))
    }

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if board[i][j] == 'O' && !visited[i][j] {
                r := region{
                    []pos{},
                    false,
                }

                DFS(i, j, board, visited, &r)

                if !r.hasPosInBorder {
                    markCells(board, r, 'X')
                }
            }
        }
    }
}

func DFS(x, y int, board [][]byte, visited [][]bool, r *region) {
    visited[x][y] = true

    (*r).positions = append((*r).positions, pos{x, y})
    if isInBorder(x, y, board) {
        (*r).hasPosInBorder = true
    }
    
    for i := 0; i < 4; i++ {
        nx := x + dx[i]
        ny := y + dy[i]

        if isIn(nx, ny, board) && board[nx][ny] == 'O' && !visited[nx][ny] {
            DFS(nx, ny, board, visited, r)
        }
    }
}

func isIn(x, y int, board [][]byte) bool {
    return 0 <= x && x < len(board) && 0 <= y && y < len(board[x])
}

func isInBorder(x, y int, board [][]byte) bool {
    return 0 == x || x == len(board) - 1 || y == 0 || y == len(board[x]) - 1
}

func markCells(board [][] byte, r region, mark byte) {
    for _, cell := range r.positions {
        x := cell.x
        y := cell.y

        board[x][y] = mark
    }
}
