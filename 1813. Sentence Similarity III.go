func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    s1 := strings.Split(sentence1, " ")
    s2 := strings.Split(sentence2, " ")

    return solve(s1, s2) || solve(s2, s1)
}

func solve(s1 []string, s2 []string) bool {
    longestPref := 0
    longestSuf := 0

    for i := 0; i < min(len(s1), len(s2)); i++ {
        if s1[i] == s2[i] {
            longestPref++
        } else {
            break
        }
    }

    for i, j := len(s1) - 1, len(s2) - 1; i >= 0 && j >= 0; i, j = i - 1, j - 1 {
        if s1[i] == s2[j] {
            longestSuf++
        } else {
            break
        }
    }

    if (longestPref + longestSuf) >= len(s1) {
        return true
    } else {
        return false
    }
}
