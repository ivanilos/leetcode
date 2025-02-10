func clearDigits(s string) string {
    st := []rune{}

    for _, c := range s {
        if unicode.IsDigit(c) {
            st = st[0 : len(st) - 1]
        } else {
            st = append(st, c)
        }
    }
    return string(st)
}
