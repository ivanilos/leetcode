func solveSudoku(board [][]byte)  {
    rowGot := make([][]int, 9)
    for i := 0; i < 9; i++ {
        rowGot[i] = make([]int, 10)
    }

    colGot := make([][]int, 9)
    for i := 0; i < 9; i++ {
        colGot[i] = make([]int, 10)
    }

    boxGot := make([][]int, 9)
    for i := 0; i < 9; i++ {
        boxGot[i] = make([]int, 10)
    }

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            boxIndex := 3 * (i / 3) + j / 3
            if board[i][j] != '.' {
                val := int(board[i][j] - '0')
                rowGot[i][val] = 1
                colGot[j][val] = 1
                boxGot[boxIndex][val] = 1
            }
        }
    }

    DFS(0, 0, board, rowGot, colGot, boxGot)
}

func DFS(curRow, curCol int, board [][]byte, rowGot, colGot, boxGot [][]int) bool {
    if curCol >= 9 {
        return DFS(curRow + 1, 0, board, rowGot, colGot, boxGot)
    }
    if curRow >= 9 {
        return true
    }

    if board[curRow][curCol] != '.' {
        return DFS(curRow, curCol + 1, board, rowGot, colGot, boxGot)
    }

    boxIndex := 3 * (curRow / 3) + curCol / 3
    for i := 1; i <= 9; i++ {
        if rowGot[curRow][i] == 0 && colGot[curCol][i] == 0 && boxGot[boxIndex][i] == 0 {
            rowGot[curRow][i] = 1 
            colGot[curCol][i] = 1
            boxGot[boxIndex][i] = 1

            board[curRow][curCol] = byte(i + int('0'))
            if DFS(curRow, curCol + 1, board, rowGot, colGot, boxGot) {
                return true
            }
            board[curRow][curCol] = '.'

            boxGot[boxIndex][i] = 0
            colGot[curCol][i] = 0
            rowGot[curRow][i] = 0
        }
    }

    return false
}
