func wordBreak(s string, wordDict []string) []string {
    wordMap := map[string]bool{}
    for _, word := range(wordDict) {
        wordMap[word] = true
    }

    return solve(0, len(s) - 1, s, wordMap)
}

func solve(l, r int, s string, wordMap map[string]bool) []string{
    result := []string{}

    for i := l; i <= r; i++ {
        piece := s[l : i + 1]
        if _, ok := wordMap[piece]; ok {
            if i == r {
                result = append(result, piece)
            } else {
                rest := solve(i + 1, r, s, wordMap)
                if len(rest) != 0 {
                    for _, val := range(rest) {
                        curSplit := piece + " " + val
                        result = append(result, curSplit)
                    }
                }
            }
        }
    }
    return result
}
