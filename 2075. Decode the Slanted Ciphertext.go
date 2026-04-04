func decodeCiphertext(encodedText string, rows int) string {
    cols := len(encodedText) / rows

    mat := make([][]byte, rows)
    for i := 0; i < rows; i++ {
        mat[i] = make([]byte, cols)
    }

    idx := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            mat[i][j] = encodedText[idx]
            idx++
        }
    }

    return getOriginalText(mat, rows, cols)
}

func getOriginalText(mat [][]byte, rows, cols int) string {
    result := []byte{}

    for startCol := 0; startCol < cols; startCol++ {
        curCol := startCol
        curRow := 0

        curRunText := []byte{}
        for i := 0; i < rows && curCol < cols; i++ {
            curRunText = append(curRunText, mat[curRow][curCol])
            curRow++
            curCol++
        }

        result = append(result, curRunText...)
    }

    for i := len(result) - 1; i >= 0; i-- {
        if result[i] != ' ' {
            result = result[: i + 1]
            break
        }
    }


    return string(result)
}
