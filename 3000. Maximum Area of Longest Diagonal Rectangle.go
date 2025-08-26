func areaOfMaxDiagonal(dimensions [][]int) int {
    maxDiagSquare := 0
    maxArea := 0

    for _, dimension := range dimensions {
        length, width := dimension[0], dimension[1]

        curDiagSquare := length * length + width * width
        if curDiagSquare > maxDiagSquare {
            maxDiagSquare = curDiagSquare
            maxArea = length * width
        } else if curDiagSquare == maxDiagSquare {
            maxArea = max(maxArea, length * width)
        }
    }

    return maxArea
}
