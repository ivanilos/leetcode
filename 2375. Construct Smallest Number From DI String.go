func smallestNumber(pattern string) string {
    used := make([]bool, 10)

    result := ""
    cur := []rune{}
    solve(-1, used, pattern, 0, cur, &result)

    return result
}

func solve(pos int, used []bool, pattern string, lastNum int, cur []rune, result *string) {
    if pos >= len(pattern) {
        candidate := string(cur)
        if *result == "" || candidate < *result {
            *result = candidate
        }
        return
    }

    for i := 1; i <= 9; i++ {
        if canAdd(i, lastNum, pos, pattern, used) {
            cur = append(cur, rune(i + '0'))
            used[i] = true

            solve(pos + 1, used, pattern, i, cur, result)

            used[i] = false
            cur = cur[: len(cur) - 1]
        }
    }
    return
}

func canAdd(num, lastNum, pos int, pattern string, used []bool) bool {
    if pos == -1 {
        return true
    }
    if used[num] {
        return false
    }

    if pattern[pos] == 'I' {
        return num > lastNum
    } else {
        return num < lastNum
    }
}
