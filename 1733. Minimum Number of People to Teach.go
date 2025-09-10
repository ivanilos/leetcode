func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
    needToTeach := map[int]bool{}

    for _, friendship := range friendships {
        a := friendship[0] - 1
        b := friendship[1] - 1
        if !knowCommonLanguage(languages[a], languages[b]) {
            needToTeach[a] = true
            needToTeach[b] = true
        }
    }

    langFreqInNeedToTeach := map[int]int{}
    maxFreqLang := 0
    for person, _ := range needToTeach {
        for _, lang := range languages[person] {
            langFreqInNeedToTeach[lang]++
            maxFreqLang = max(maxFreqLang, langFreqInNeedToTeach[lang])
        }
    }

    return len(needToTeach) - maxFreqLang
}

func knowCommonLanguage(langA, langB []int) bool {
    mapLangA := map[int]bool{}
    for _, l := range langA {
        mapLangA[l] = true
    }

    for _, l := range langB {
        if _, ok := mapLangA[l]; ok {
            return true
        }
    }

    return false
}
