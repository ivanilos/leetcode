func getHappyString(n int, k int) string {
    result := ""
    DFS(n, &k, -1, []rune{}, &result)

    return result
}

func DFS(left int, k *int, lastCharIdx int, curString []rune, result *string) {
    if baseConditionReached(left, k, curString, result) {
        return
    }

    for i := 0; i < 3; i++ {
        if i == lastCharIdx {
            continue
        }
        curChar := rune(i + int('a'))
        curString = append(curString, curChar)

        DFS(left - 1, k, i, curString, result)

        curString = curString[: len(curString) - 1]
    }
}

func baseConditionReached(left int, k *int, curString []rune, result *string) bool {
    if *k <= 0 {
        return true
    }

    if left == 0 {
        *k--
        if *k == 0 {
            *result = string(curString)
        }
        return true
    }

    return false
}
