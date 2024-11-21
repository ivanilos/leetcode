type Node struct {
    Pos
    dir int
}

type Pos struct {
    x, y int
}

const VER int = 0
const HOR int = 1

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    usedCells := map[Node]bool{}

    guardsMap := getGuardsMap(guards)
    wallsMap := getWallsMap(walls)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if _, ok := guardsMap[Pos{i, j}]; ok {
                markGuarded(usedCells, guardsMap, wallsMap, i, j, n, m)
            }
        }
    }

    guarded := map[Pos]bool{}
    for cell, _ := range usedCells {
        guarded[Pos{cell.x, cell.y}] = true
    }

    return m * n - len(guarded) - len(guardsMap) - len(wallsMap)
}

func getGuardsMap(guards [][]int) map[Pos]bool {
    guardsMap := map[Pos]bool{}

    for _, guard := range guards {
        guardsMap[Pos{guard[0], guard[1]}] = true
    }
    return guardsMap
}

func getWallsMap(walls [][]int) map[Pos]bool {
    wallsMap := map[Pos]bool{}

    for _, wall := range walls {
        wallsMap[Pos{wall[0], wall[1]}] = true
    }
    return wallsMap
}

func markGuarded(usedCells map[Node]bool, guardsMap, wallsMap map[Pos]bool, x, y, n, m int) {

    for dir := -1; dir <= 1; dir += 2 {
        // vertical movement
        for step := 1; step < m; step++ {
            nx := x + step * dir

            if isValidUnvisitedPos(nx, y, VER, n, m, guardsMap,wallsMap,usedCells) && markValidPos(nx, y, VER, usedCells) {
                ;
            } else {
                break
            }
        }

        for step := 1; step < n; step++ {
            ny := y + step * dir

            if isValidUnvisitedPos(x, ny, HOR, n, m, guardsMap,wallsMap,usedCells) && markValidPos(x, ny, HOR, usedCells) {
                ;
            } else {
                break
            }
        }
    }
    return
}

func markValidPos(x, y, dir int, usedCells map[Node]bool) bool {
    usedCells[Node{Pos{x, y}, dir}] = true
    return true
}

func isValidUnvisitedPos(x, y, dir, n, m int, guardsMap, wallsMap map[Pos]bool, usedCells map[Node]bool) bool {
    if x < 0 || x >= m {
        return false
    }
    if y < 0 || y >= n {
        return false
    }

    pos := Pos{x, y}
    if _, ok := usedCells[Node{pos, dir}]; ok {
        return false
    }
    if _, ok := guardsMap[pos]; ok {
        return false
    }
    if _, ok := wallsMap[pos]; ok {
        return false
    }

    return true
}
