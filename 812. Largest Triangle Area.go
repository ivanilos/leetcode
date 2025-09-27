func largestTriangleArea(points [][]int) float64 {
    result := float64(0)

    for a := 0; a < len(points); a++ {
        for b := a + 1; b < len(points); b++ {
            for c := b + 1; c < len(points); c++ {
                result = max(result, area(points[a], points[b], points[c]) / 2)
            }
        }
    }

    return result
}

func area(a, b, c []int) float64 {
    return float64(abs(a[0] * (b[1] - c[1]) + b[0] * (c[1] - a[1]) + c[0] * (a[1] - b[1])))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
