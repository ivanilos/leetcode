func numberOfPairs(points [][]int) int {
    result := 0

    for upperLeftIdx := 0; upperLeftIdx < len(points); upperLeftIdx++ {
        for bottomRightIdx := 0; bottomRightIdx < len(points); bottomRightIdx++ {
            if upperLeftIdx == bottomRightIdx {
                continue
            }
            if !rectangleCoordinates(points[upperLeftIdx], points[bottomRightIdx]) {
                continue
            }

            goodRectangle := true
            for i, point := range points {
                if i == upperLeftIdx || i == bottomRightIdx {
                    continue
                }

                if inside(points[upperLeftIdx], points[bottomRightIdx], point) {
                    goodRectangle = false
                    break
                }
            }

            if goodRectangle {
                result++
            }
        }
    }

    return result
}

func rectangleCoordinates(A, B []int) bool {
    return A[0] <= B[0] && A[1] >= B[1]
}

func inside(A, B, C []int) bool {
    return A[0] <= C[0] && C[0] <= B[0] && A[1] >= C[1] && C[1] >= B[1]
}
