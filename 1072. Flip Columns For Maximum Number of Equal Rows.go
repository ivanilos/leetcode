func maxEqualRowsAfterFlips(matrix [][]int) int {
    freq := map[string]int{}
    for _, row := range matrix {
        rowVal := getRowVal(row)
        rowValInv := getRowValInv(row)

        freq[rowVal]++
        freq[rowValInv]++
    }

    result := 0
    for _, val := range freq {
        result = max(result, val)
    }
    return result
}

func getRowVal(row []int) string {
    result := make([]int, len(row))

    for i := 0; i < len(row); i++ {
        result[i] = row[i]
    }

    return strings.Trim(strings.Replace(fmt.Sprint(result), " ", "", -1), "[]")
}

func getRowValInv(row []int) string {
    result := make([]int, len(row))

    for i := 0; i < len(row); i++ {
        result[i] = 1 - row[i]
    }

    return strings.Trim(strings.Replace(fmt.Sprint(result), " ", "", -1), "[]")
}
