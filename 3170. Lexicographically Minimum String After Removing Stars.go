func clearStars(s string) string {
    const INF = int(1e9)

    st := make([][]int, 26)
    for i := 0; i < 26; i++ {
        st[i] = []int{}
    }

    for i, c := range s {
        if c == '*' {
            for j := 0; j < 26; j++ {
                if len(st[j]) > 0 {
                    st[j] = st[j][: len(st[j]) - 1]
                    break
                }
            }
        } else {
            idx := int(c - 'a')
            st[idx] = append(st[idx], i)
        }
    }

    result := []rune{}
    next := make([]int, 26)
    changed := true

    for changed == true {
        changed = false

        mini := INF
        minIdx := 0
        for i := 0; i < 26; i++ {
            if next[i] < len(st[i]) && st[i][next[i]] < mini {
                mini = st[i][next[i]]
                minIdx = i
                changed = true
            }
        }

        if changed {
            next[minIdx]++
            nextChar := rune(minIdx + 'a')
            result = append(result, nextChar)
        }
    }

    return string(result)
}
