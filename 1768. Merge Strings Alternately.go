func mergeAlternately(word1 string, word2 string) string {
    result := []byte{}

    for i := 0; i < min(len(word1), len(word2)); i++ {
        result = append(result, word1[i])
        result = append(result, word2[i])
    }

    for i := min(len(word1), len(word2)); i < len(word1); i++ {
        result = append(result, word1[i])
    }

    for i := min(len(word1), len(word2)); i < len(word2); i++ {
        result = append(result, word2[i])
    }

    return string(result)
}
