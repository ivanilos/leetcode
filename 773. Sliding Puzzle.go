const ROWS int = 2
const COLS int = 3

func slidingPuzzle(board [][]int) int {
    result := BFS(board)
    return result
}

func BFS(board [][]int) int {
    dist := map[int]int{}

    start := encode(board)
    goal := encode([][]int{{1, 2, 3}, {4, 5, 0}})
    dist[start] = 0

    q := []int{start}

    for len(q) > 0 {
        encodedCurBoard := q[0]
        curBoard := decode(encodedCurBoard)
        q = q[1 : len(q)]

        encodedNextPositions := getEncodedNextPositions(curBoard)

        for _, encodedNextPosition := range encodedNextPositions {
            if _, ok := dist[encodedNextPosition]; !ok {
                dist[encodedNextPosition] = 1 + dist[encodedCurBoard]
                q = append(q, encodedNextPosition)
            }
        }
    }

    if _, ok := dist[goal]; !ok {
        return -1
    }

    return dist[goal]
}

func encode(board [][]int) int {    
    result := 0
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            result = result * 10 + board[i][j]
        }
    }
    return result
}

func decode(val int) [][]int {
    result := make([][]int, ROWS)
    for i := 0; i < ROWS; i++ {
        result[i] = make([]int, COLS)
    }

    for i := ROWS - 1; i >= 0; i-- {
        for j := COLS - 1; j >= 0; j-- {
            result[i][j] = val % 10
            val /= 10
        }
    }
    return result
}

func getEncodedNextPositions(board [][]int) []int {
    zeroX, zeroY := getZeroPos(board)

    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}

    result := []int{}
    for i := 0; i < 4; i++ {
        nx := zeroX + dx[i]
        ny := zeroY + dy[i]

        if isIn(nx, ny) {
            board[nx][ny], board[zeroX][zeroY] = board[zeroX][zeroY], board[nx][ny]
            result = append(result, encode(board))
            board[nx][ny], board[zeroX][zeroY] = board[zeroX][zeroY], board[nx][ny]
        }
    }

    return result
}

func isIn(nx, ny int) bool {
    return 0 <= nx && nx < ROWS && 0 <= ny && ny < COLS
}

func getZeroPos(board [][]int) (int, int) {
    zeroX := 0
    zeroY := 0

    for i := 0; i < ROWS; i++ {
        for j := 0; j < COLS; j++ {
            if board[i][j] == 0 {
                zeroX = i
                zeroY = j
                i = 2
                break
            }
        }
    }

    return zeroX, zeroY
}
