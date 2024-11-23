func rotateTheBox(box [][]byte) [][]byte {
    boxWithPositionsChanged := applyGravity(box)
    result := rotateClockwise(boxWithPositionsChanged)

    return result
}

func applyGravity(box [][]byte) [][]byte {
    for i := 0; i < len(box); i++ {
        nextFreePos := len(box[i]) - 1
        for j := len(box[i]) - 1; j >= 0; j-- {
            if box[i][j] == '#' {
                box[i][j] = '.'
                box[i][nextFreePos] = '#'
                nextFreePos--
            } else if box[i][j] == '*' {
                nextFreePos = j - 1
            }
        }
    }
    return box
}

func rotateClockwise(box [][]byte) [][]byte {
    result := make([][]byte, len(box[0]))

    for i := 0; i < len(box[0]); i++ {
        result[i] = make([]byte, len(box))

        for j := 0; j < len(box); j++ {
            result[i][j] = box[len(box) - 1 - j][i]
        }
    }
    return result
}
