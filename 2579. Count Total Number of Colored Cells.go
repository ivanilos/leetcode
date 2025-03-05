func coloredCells(n int) int64 {
    result := int64(2 * n - 1)

    for i := int64(2); term(int64(n), i) > 0 ; i++ {
        result += term(int64(n), i)
    }

    return result
}

func term(n, i int64) int64 {
    return int64(2) * (int64(2) * n - (int64(2) * i - int64(1)))
}
