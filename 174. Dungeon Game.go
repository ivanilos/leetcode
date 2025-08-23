const INF = int(1e9)

func calculateMinimumHP(dungeon [][]int) int {
    n := len(dungeon)
    m := len(dungeon[0])

    low := 1
    high := n * m * 1000
    result := high

    for low <= high {
        mid := (low + high) / 2

        if canTraverse(mid, dungeon, n, m) {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return result
}

func canTraverse(startEnergy int, dungeon [][]int, n, m int) bool {
    energy := make([][]int, n)
    for i := 0; i < n; i++ {
        energy[i] = make([]int, m)
    }

    energy[0][0] = startEnergy + dungeon[0][0]

    for j := 1; j < m; j++ {
        if energy[0][j - 1] > 0 {
            energy[0][j] = energy[0][j - 1] + dungeon[0][j]
        }
    }

    for i := 1; i < n; i++ {
        if energy[i - 1][0] > 0 {
            energy[i][0] = energy[i - 1][0] + dungeon[i][0]
        }
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            energy[i][j] = -1
            if energy[i][j - 1] > 0 {
                energy[i][j] = max(energy[i][j], energy[i][j - 1] + dungeon[i][j])
            }
            if energy[i - 1][j] > 0 {
                energy[i][j] = max(energy[i][j], energy[i - 1][j] + dungeon[i][j])
            }
        }
    }
    return energy[n - 1][m - 1] > 0
}
