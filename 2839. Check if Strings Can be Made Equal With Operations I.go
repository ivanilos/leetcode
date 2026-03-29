func canBeEqual(s1 string, s2 string) bool {
    a1 := []byte(s1)
    a2 := []byte(s2)

    makeCanonical(a1)
    makeCanonical(a2)

    return string(a1) == string(a2)
}

func makeCanonical(s []byte) {
    if s[0] > s[2] {
        s[0], s[2] = s[2], s[0]
    }
    if s[1] > s[3] {
        s[1], s[3] = s[3], s[1]
    }
}
