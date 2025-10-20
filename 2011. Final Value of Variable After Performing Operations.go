func finalValueAfterOperations(operations []string) int {
    result := 0
    for _, c := range operations {
        if c[0] == '-' {
            result--
        } else if c[2] == '-' {
            result--
        } else {
            result++
        }
    }
    return result
}
