func reverseWords(s string) string {
    words := strings.Split(s, " ")

    result := []string{}
    for i := len(words) - 1; i >= 0; i-- {
        if len(words[i]) != 0 {
            result = append(result, words[i])
        }
    }

    return strings.Join(result, " ")
}
