func maxPoints(points [][]int) int {
    if len(points) <= 1 {
        return len(points)
    }
    
    result := 0
    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            curLinePoints := 0
            for k := 0; k < len(points); k++ {
                if isInline(points[i], points[j], points[k]) {
                    curLinePoints++
                }
            }
            result = max(result, curLinePoints)
        }
    }
    return result
}

func isInline(a []int, b []int, c []int) bool {
    return a[0] * (b[1] - c[1]) + b[0] * (c[1] - a[1]) + c[0] * (a[1] - b[1]) == 0
}
