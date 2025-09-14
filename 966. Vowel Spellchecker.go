var vowels = map[byte]bool {
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
}

func spellchecker(wordlist []string, queries []string) []string {
    result := make([]string, len(queries))

    wordsMap := map[string]bool{}
    wordsInLowercaseMap := map[string]string{}
    wordsDevoweledMap := map[string]string{}

    for _, word := range wordlist {
        wordsMap[word] = true

        lowercaseWord := strings.ToLower(word)
        if _, ok := wordsInLowercaseMap[lowercaseWord]; !ok {
            wordsInLowercaseMap[lowercaseWord] = word
        }

        devoweledWord := devowel(word)
        if _, ok := wordsDevoweledMap[devoweledWord]; !ok {
            wordsDevoweledMap[devoweledWord] = word
        }
    }

    for i, query := range queries {
        result[i] = match(wordsMap, wordsInLowercaseMap, wordsDevoweledMap, wordlist, query)
    }

    return result
}

func match(wordsMap map[string]bool, wordsInLowercaseMap, wordsDevoweledMap map[string]string, wordlist[]string, query string) string {
    if _, ok := wordsMap[query]; ok {
        return query
    }

    if _, ok := wordsInLowercaseMap[strings.ToLower(query)]; ok {
        return wordsInLowercaseMap[strings.ToLower(query)]
    }

    devoweledQuery := devowel(query)
    if _, ok := wordsDevoweledMap[devoweledQuery]; ok {
        return wordsDevoweledMap[devoweledQuery]
    }

    return ""
}

func devowel(s string) string {
    str := []byte(strings.ToLower(s))

    for i := 0; i < len(str); i++ {
        if isVowel(str[i]) {
            str[i] = 'a'
        }
    }

    return string(str)
}

func isVowel(ch byte) bool {
    _, ok := vowels[ch]

    return ok
}
