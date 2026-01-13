const MAX_IT = 100

func separateSquares(squares [][]int) float64 {
    totalArea := float64(0)

    for _, square := range squares {
        totalArea += float64(square[2] * square[2])
    }

    low := float64(0)
    high := float64(2e9)
    best := float64(0)

    for it := 0; it < MAX_IT; it++ {
        mid := (low + high) / 2

        if getAreaBelowLine(squares, mid) >= totalArea / 2 {
            high = mid
        } else {
            best = mid
            low = mid
        }
    }

    return best
}

func getAreaBelowLine(squares [][]int, yLine float64) float64 {
    result := float64(0)

    for _, square := range squares {
        side := float64(square[2])
        y := float64(square[1])
        result += side * max(0.0, min(side, yLine - y))
    }

    return result
}
