func firstCompleteIndex(arr []int, mat [][]int) int {
    numToRow := map[int]int{}
    numToCol := map[int]int{}

    rows := make([]int, len(mat))
    cols := make([]int, len(mat[0]))

    for i := 0; i < len(mat);i++ {
        for j := 0; j < len(mat[i]); j++ {
            numToRow[mat[i][j]] = i
            numToCol[mat[i][j]] = j
        }
    }

    for idx, val := range arr {
        rows[numToRow[val]]++
        cols[numToCol[val]]++

        if rows[numToRow[val]] == len(mat[0]) || cols[numToCol[val]] == len(mat) {
            return idx
        }
    }
    panic("Should not reach here")
}
