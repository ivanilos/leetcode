var dx []int = []int{1, 0, -1, 0}
var dy []int = []int{0, 1, 0, -1}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    islands := getIslands(grid2)

    result := 0
    for _, island := range islands {
        if isContained(island, grid1) {
            result++
        }
    }
    return result
}

func getIslands(grid [][]int) [][][]int {
    visited := make([][]bool, len(grid))
    for i := 0; i < len(visited); i++ {
        visited[i] = make([]bool, len(grid[i]))
    }

    result := [][][]int{}
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if !visited[i][j] && grid[i][j] == 1 {
                curIsland := [][]int{}
                DFS(i, j, grid, visited, &curIsland)

                result = append(result, curIsland)
            }
        }
    }
    return result
}

func DFS(i, j int, grid [][]int, visited [][]bool, curIsland *[][]int) {
    visited[i][j] = true
    *curIsland = append(*curIsland, []int{i, j})

    for k := 0; k < 4; k++ {
        x := i + dx[k]
        y := j + dy[k]

        if isIn(x, y, grid) && grid[x][y] == 1 && !visited[x][y] {
            DFS(x, y, grid, visited, curIsland)
        }
    }
}

func isIn(x, y int, grid [][]int) bool {
    return 0 <= x && x < len(grid) && 0 <= y && y < len(grid[x])
}

func isContained(island [][]int, grid [][]int) bool {
    for i := 0; i < len(island); i++ {
        x := island[i][0]
        y := island[i][1]
        if grid[x][y] != 1 {
            return false
        }
    }
    return true
}
