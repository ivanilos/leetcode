const MOD = int64(1e9 + 7)

func countTrapezoids(points [][]int) int {
    yToXFreq := map[int]int64{}

    for _, point := range points {
        yToXFreq[point[1]]++
    }

    result := int64(0)
    parallelSides := int64(0)

    for _, qt := range yToXFreq {
        nextParallelQt := (qt * (qt - 1)) / 2
        result = (result + parallelSides * nextParallelQt) % MOD

        parallelSides += nextParallelQt
    }

    return int(result)
}
