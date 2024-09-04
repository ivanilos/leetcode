type Pair struct {
    x, y int
}

func robotSim(commands []int, obstacles [][]int) int {
    obstacleMap := map[Pair]bool{}

    for _, obstacle := range obstacles {
        obstacleMap[Pair{obstacle[0], obstacle[1]}] = true
    }

    x := 0
    y := 0
    dir := 0
    dx := []int{0, 1, 0, -1}
    dy := []int{1, 0, -1, 0}
    N_DIRS := 4

    result := 0
    for _, command := range commands {
        if command > 0 {
            for i := 0; i < command; i++ {
                nx := x + dx[dir]
                ny := y + dy[dir]

                if _, ok := obstacleMap[Pair{nx, ny}]; !ok {
                    x = nx
                    y = ny
                }
                result = max(result, x * x + y * y)
            }
        } else if command == -1 {
            dir = (dir + 1) % N_DIRS
        } else {
            dir = (dir - 1 + N_DIRS) % N_DIRS
        }
    }
    return result
}
