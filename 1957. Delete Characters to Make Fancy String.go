func makeFancyString(s string) string {
    result := []rune{}

    lastTwo := make([]rune, 2)
    for _, c := range s {
        if lastTwo[0] == lastTwo[1] && lastTwo[1] == c {
            continue
        }

        result = append(result, c)
        lastTwo[0] = lastTwo[1]
        lastTwo[1] = c
    }

    return string(result)
}
