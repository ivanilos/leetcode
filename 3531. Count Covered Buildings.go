const INF = int(1e9)

func countCoveredBuildings(n int, buildings [][]int) int {
    xToYMin := map[int]int{}
    xToYMax := map[int]int{}
    yToXMin := map[int]int{}
    yToXMax := map[int]int{}

    for _, building := range buildings {
        x := building[0]
        y := building[1]

        xToYMin[x] = INF
        xToYMax[x] = -INF
        yToXMin[y] = INF
        yToXMax[y] = -INF
    }

    for _, building := range buildings {
        x := building[0]
        y := building[1]

        xToYMin[x] = min(xToYMin[x], y)
        xToYMax[x] = max(xToYMax[x], y)
        yToXMin[y] = min(yToXMin[y], x)
        yToXMax[y] = max(yToXMax[y], x)
    }

    result := 0
    for _, building := range buildings {
        x := building[0]
        y := building[1]

        if xToYMin[x] < y && y < xToYMax[x] && yToXMin[y] < x && x <  yToXMax[y] {
            result++
        }
    }

    return result
}
