func removeOccurrences(s string, part string) string {
    result := []rune{}

    for _, c := range s {
        result = append(result, c)

        for len(result) >= len(part) && string(result[len(result) - len(part) : len(result)]) == part {
            result = result[0 : len(result) - len(part)]
        }
    }

    return string(result)
}
