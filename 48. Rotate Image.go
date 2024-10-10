func rotate(matrix [][]int)  {
    _rotate(matrix, 0, 0, len(matrix) - 1, len(matrix) - 1)
}

func _rotate(matrix [][]int, rb, cb, re, ce int) {
    if rb >= re && cb >= ce {
        return
    }

    for j := cb; j < ce; j++ {
        nextVal := matrix[rb][j]

        posRow := rb
        posCol := j
        for k := 0; k < 4; k++ {
            nextPosRow := posCol
            nextPosCol := rb + (re - posRow)

            temp := matrix[nextPosRow][nextPosCol]
            matrix[nextPosRow][nextPosCol] = nextVal
            nextVal = temp
            
            posRow = nextPosRow
            posCol = nextPosCol
        }
    }

    _rotate(matrix, rb + 1, cb + 1, re - 1, ce - 1)
}
