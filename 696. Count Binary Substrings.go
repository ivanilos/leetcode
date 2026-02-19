func countBinarySubstrings(s string) int {
    splittedQuantity := []int{}

    last := s[0]
    qt := 0
    for i := 0; i < len(s); i++ {
        if s[i] == last {
            qt++
        } else {
            splittedQuantity = append(splittedQuantity, qt)
            qt = 1
        }
        last = s[i]
    }

    splittedQuantity = append(splittedQuantity, qt)

    result := 0
    for i := 1; i < len(splittedQuantity); i++ {
        result += min(splittedQuantity[i], splittedQuantity[i - 1])
    }

    return result
}
