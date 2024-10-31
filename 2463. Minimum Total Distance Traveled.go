const INF int64 = int64(1e17)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
    sort.Slice(robot, func(a int, b int) bool {
        return robot[a] < robot[b]
    })
    sort.Slice(factory, func(a int, b int) bool {
        return factory[a][0] < factory[b][0]
    })

    dp := make([][][]int64, len(robot))
    for i := 0; i < len(robot); i++ {
        dp[i] = make([][]int64, len(factory))
        for j := 0; j < len(factory); j++ {
            dp[i][j] = make([]int64, factory[j][1] + 1)
            for k := 0; k < factory[j][1] + 1; k++ {
                dp[i][j][k] = -1
            }
        }
    }

    result := solve(0, 0, 0, robot, factory, dp)
    return result
}

func solve(nextRobot, nextFactory, usedLimit int, robot []int, factory [][]int, dp [][][]int64) int64 {
    if nextRobot >= len(robot) {
        return 0
    }
    if nextFactory >= len(factory) {
        return INF
    }
    if usedLimit >= factory[nextFactory][1] {
        return solve(nextRobot, nextFactory + 1, 0, robot, factory, dp)
    }
    if dp[nextRobot][nextFactory][usedLimit] != -1 {
        return dp[nextRobot][nextFactory][usedLimit]
    }

    // skip current factory
    result := solve(nextRobot, nextFactory + 1, 0, robot, factory, dp)

    // use current factory
    dist := getDist(int64(robot[nextRobot]), int64(factory[nextFactory][0]))
    result = min(result, dist + solve(nextRobot + 1, nextFactory, usedLimit + 1, robot, factory, dp))

    dp[nextRobot][nextFactory][usedLimit] = result
    return result
}

func getDist(x int64, y int64) int64 {
    return abs(x - y)
}

func abs(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x
}
