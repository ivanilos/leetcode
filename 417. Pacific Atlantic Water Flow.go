var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func pacificAtlantic(heights [][]int) [][]int {
    canReachPacific := getReachabilityMatrix(0, 0, heights)
    canReachAtlantic := getReachabilityMatrix(len(heights) - 1, len(heights[0]) - 1, heights)

    result := [][]int{}
    for i := 0; i < len(heights); i++ {
        for j := 0; j < len(heights[i]); j++ {
            if canReachPacific[i][j] && canReachAtlantic[i][j] {
                result = append(result, []int{i, j})
            }
        }
    }

    return result
}

func newReachMatrix(heights [][]int) [][]bool {
    reachOcean := make([][]bool, len(heights))
    for i := 0; i < len(heights); i++ {
        reachOcean[i] = make([]bool, len(heights[i]))
    }

    return reachOcean
}

func getReachabilityMatrix(rowAdjacent, colAdjacent int, heights [][]int) [][]bool {
    rows := len(heights)
    cols := len(heights[0])

    canReachOcean := newReachMatrix(heights)

    q := [][]int{}
    for j := 0; j < cols; j++ {
        canReachOcean[rowAdjacent][j] = true
        q = append(q, []int{rowAdjacent, j})
    }

    for i := 0; i < rows; i++ {
        canReachOcean[i][colAdjacent] = true
        q = append(q, []int{i, colAdjacent})
    }

    for len(q) > 0 {
        x := q[0][0]
        y := q[0][1]

        q = q[1 : len(q)]

        for i := 0; i < 4; i++ {
            nx := x + dx[i]
            ny := y + dy[i]

            if isIn(nx, ny, rows, cols) && heights[nx][ny] >= heights[x][y] && !canReachOcean[nx][ny] {
                canReachOcean[nx][ny] = true
                q = append(q, []int{nx, ny})
            }
        }
    }

    return canReachOcean
}

func isIn(x, y, rows, cols int) bool {
    return 0 <= x && x < rows && 0 <= y && y < cols
}
