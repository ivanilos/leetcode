func longestValidParentheses(s string) int {
    result := 0

    prevPos := map[int]int{}
    prevPos[0] = -1
    curVal := 0
    st := [][]int{[]int{-1, 0}}
    for i, ch := range s {
        if ch == '(' {
            curVal++
            if _, ok := prevPos[curVal]; !ok {
                prevPos[curVal] = i
            }
            st = append(st, []int{i, curVal})
        } else {
            curVal--
            for len(st) > 0 && st[len(st) - 1][1] > curVal {
                delete(prevPos, st[len(st) - 1][1])
                st = st[: len(st) - 1]
            }

            if _, ok := prevPos[curVal]; !ok {
                prevPos[curVal] = i
            }

            st = append(st, []int{i, curVal})
        }

        if _, ok := prevPos[curVal]; ok {
            result = max(result, i - prevPos[curVal])
        }
    }

    return result
}
