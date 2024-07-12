func maximumGain(s string, x int, y int) int {
    result := 0
    if (x > y) {
        v1, str := solve(s, x, []rune{'a', 'b'})
        v2, _ := solve(str, y, []rune{'b', 'a'})
        result = v1 + v2
    } else {
        v1, str := solve(s, y, []rune{'b', 'a'})
        v2, _ := solve(str, x, []rune{'a', 'b'})
        result = v1 + v2
    }
    return result
}

func solve(s string, val int, pattern []rune) (int, string) {
    st := []rune{}

    result := 0
    for _, ch := range(s) {
        if len(st) > 0 && st[len(st) - 1] == pattern[0] && ch == pattern[1] {
            result += val
            st = st[0 : len(st) - 1]
        } else {
            st = append(st, ch)
        }
    }
    return result, string(st)
}
