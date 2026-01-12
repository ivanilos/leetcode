func minTimeToVisitAllPoints(points [][]int) int {
    result := 0

    for i := 1; i < len(points); i++ {
        result += max(abs(points[i][0] - points[i - 1][0]),
                        abs(points[i][1] - points[i - 1][1]))
    }

    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
