func addSpaces(s string, spaces []int) string {
    result := []byte{}

    nextSpaceIdx := 0
    spaces = append(spaces, len(s))

    for i := 0; i < len(s); i++ {
        if spaces[nextSpaceIdx] == i {
            result = append(result, ' ')
            nextSpaceIdx++
        }
        result = append(result, s[i])
    }

    return string(result)
}
