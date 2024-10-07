func minLength(s string) int {
    st := []rune{}

    for _, c := range s {
        st = append(st, c)

        for canRemove(st, 'A', 'B') || canRemove(st, 'C', 'D') {
            st = st[0 : len(st) - 2]
        }
    }
    return len(st)
}

func canRemove(st []rune, first, sec rune) bool {
    sz := len(st)

    return sz >= 2 && st[sz - 2] == first && st[sz - 1] == sec
}
