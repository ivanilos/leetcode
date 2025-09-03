const INF = int(1e9 + 1)

func numberOfPairs(points [][]int) int {
    slices.SortFunc(points, func(a, b []int) int {
        if a[0] == b[0] {
            return b[1] - a[1]
        }
        return a[0] - b[0]
    })

    result := 0
    for i := 0; i < len(points); i++ {
        A := points[i]

        minY := -INF
        minX := A[0] - 1
        maxY := A[1] + 1
        for j := i + 1; j < len(points); j++ {
            B := points[j]

            if B[0] > minX && B[1] > minY && B[1] < maxY {
                result++
                minY = B[1]
                minX = B[0]
            }
        }
    }

    return result
}
