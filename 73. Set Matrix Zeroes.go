func setZeroes(matrix [][]int)  {
    changeFirstRow := false

    for j := 0; j < len(matrix[0]); j++ {
        if matrix[0][j] == 0 {
            changeFirstRow = true
            break
        }
    }

    for i := 1; i < len(matrix); i++ {

        changeCurRow := false;
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == 0 {
                changeCurRow = true
                break
            }
        }

        if changeCurRow {
            for j := 0; j < len(matrix[i]); j++ {
                if matrix[i][j] == 0 {
                    matrix[0][j] = 0
                }
                matrix[i][j] = 0
            }
        }
    }

    for j := 0; j < len(matrix[0]); j++ {
        if matrix[0][j] == 0 {
            for i := 1; i < len(matrix); i++ {
                matrix[i][j] = 0
            }
        }

        if changeFirstRow {
            matrix[0][j] = 0
        }
    }
}
