func compressedString(word string) string {
    result := []string{}

    lastChar := rune(word[0])
    cur := 0
    for _, c := range word {
        if c == lastChar && cur < 9 {
            cur++
        } else {
            result = append(result, strconv.Itoa(cur), string(lastChar))
            cur = 1
            lastChar = c
        }
    }

    if cur > 0 {
        result = append(result, strconv.Itoa(cur), string(lastChar))
    }
    return  strings.Join(result, "")
}
