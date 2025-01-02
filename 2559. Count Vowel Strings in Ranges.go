func vowelStrings(words []string, queries [][]int) []int {
    pref := make([]int, len(words))

    for i := 0; i < len(words); i++ {
        pref[i] = startAndEndWithVowel(words[i])

        if i > 0 {
            pref[i] += pref[i - 1]
        }
    }

    fmt.Println(pref)

    result := make([]int, len(queries))
    for i := 0; i < len(queries); i++ {
        result[i] = pref[queries[i][1]]

        if queries[i][0] - 1 >= 0 {
            result[i] -= pref[queries[i][0] - 1]
        }
    }

    return result
}

func startAndEndWithVowel(s string) int {
    last := len(s) - 1

    if isVowel(s[0]) && isVowel(s[last]) {
        return 1
    } else {
        return 0
    }
}

func isVowel(c byte) bool {
    vowels := []byte{'a', 'e', 'i', 'o', 'u'}

    for _, v := range vowels {
        if c == v {
            return true
        }
    }
    return false
}
