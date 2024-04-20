func findFarmland(land [][]int) [][]int {
    used := make([][]bool, len(land))
    for i := 0; i < len(land); i++ {
        used[i] = make([]bool, len(land[i])) 
    }

    result := [][]int{}
    for i := 0; i < len(land); i++ {
        for j := 0; j < len(land[i]); j++ {
            if !used[i][j] && land[i][j] == 1 {
                result = append(result, getLand(i, j, land, used))
            }
        }
    }
    return result
}

func getLand(i, j int, land [][]int, used [][]bool) []int {
    result := []int{i, j}

    lastX := i
    lastY := j
    for x := i; x < len(land); x++ {
        for y := j; y < len(land[x]); y++ {
            if land[x][y] == 1 {
                used[x][y] = true
                lastX = x
                lastY = y
            } else {
                if y == j {
                    x = len(land)
                }
                break
            }
        }
    }
    result = append(result, lastX)
    result = append(result, lastY)
    return result
}
