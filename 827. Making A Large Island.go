var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func largestIsland(grid [][]int) int {
    used, landSize, landIdx := pre(grid)

    landCount := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 && used[i][j] == 0 {
                landTiles := [][]int{}
                sz := DFS(i, j, grid, used, &landTiles)

                for _, tile := range landTiles {
                    landIdx[tile[0]][tile[1]] = landCount
                    landSize[tile[0]][tile[1]] = sz
                }
                landCount++
            }
        }
    }

    return solve(grid, landSize, landIdx)
}

func solve(grid, landSize, landIdx [][]int) int {
    result := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                result = max(result, landSize[i][j])
            } else {

                usedLandIdx := map[int]bool{}

                curSize := 1
                for k := 0; k < 4; k++ {
                    nx := i + dx[k]
                    ny := j + dy[k]

                    if isIn(nx, ny, grid) && grid[nx][ny] == 1 {
                        if _, ok := usedLandIdx[landIdx[nx][ny]]; !ok {
                            curSize += landSize[nx][ny]
                            usedLandIdx[landIdx[nx][ny]] = true
                        }
                    }
                }

                result = max(result, curSize)
            }
        }
    }
    return result
}

func DFS(x, y int, grid, used [][]int, landTiles *[][]int) int {
    used[x][y] = 1

    *landTiles = append(*landTiles, []int{x, y})

    sz := 1
    for i := 0; i < 4; i++ {
        nx := x + dx[i]
        ny := y + dy[i]

        if isIn(nx, ny, grid) && used[nx][ny] == 0 && grid[nx][ny] == 1 {
            sz += DFS(nx, ny, grid, used, landTiles)
        }
    }
    return sz
}

func isIn(x, y int, grid [][]int) bool {
    return 0 <= x && x < len(grid) && 0 <= y && y < len(grid[x])
}

func pre(grid [][]int) ([][]int, [][]int, [][]int) {
    used := make([][]int, len(grid))
    for i := 0; i < len(grid); i++ {
        used[i] = make([]int, len(grid[i]))
    }

    landSize := make([][]int, len(grid))
    landIdx := make([][]int, len(grid))
    for i := 0; i < len(grid); i++ {
        landSize[i] = make([]int, len(grid[i]))
        landIdx[i] = make([]int, len(grid[i]))
    }

    return used, landSize, landIdx
}
