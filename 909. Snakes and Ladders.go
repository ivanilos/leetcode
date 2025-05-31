var INF = int(1e9)

func snakesAndLadders(board [][]int) int {
    n := len(board)

    result := BFS(1, n * n, n, board)

    return result
}

func BFS(start, end, n int, board [][]int) int {
    dist := make([]int, end + 1)
    for i := start; i <= end; i++ {
        dist[i] = INF
    }

    q := []int{start}
    dist[start] = 0

    for len(q) > 0 {
        cur := q[0]
        q = q[1 : len(q)]

        for next := cur + 1; next <= cur + 6 && next <= end; next++ {
            boardX, boardY := getBoardPos(next, n)

            nextPos := next
            if board[boardX][boardY] != -1 {
                nextPos = board[boardX][boardY]
            }

            if dist[nextPos] > 1 + dist[cur] {
                dist[nextPos] = 1 + dist[cur]
                q = append(q, nextPos)
            }
        }
    }

    if dist[end] < INF {
        return dist[end]
    }
    return -1
}

func getBoardPos(pos, n int) (int, int) {
    pos--

    row := n - 1 - (pos / n)
    col := -1

    if (n - 1 - row) % 2 == 0 {
        col = pos - (n - 1 - row) * n
    } else {
        col = n - 1 - (pos - (n - 1 - row) * n)
    }
        
    return row, col
}
