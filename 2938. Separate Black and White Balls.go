func minimumSteps(s string) int64 {
    result := int64(0)

    curBlack := int64(0)
    for _, c := range s {
        if c == '1' {
            curBlack++
        } else {
            result += curBlack
        }
    }
    return result
}
