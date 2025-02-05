func areAlmostEqual(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }

    diff1 := []byte{}
    diff2 := []byte{}
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            diff1 = append(diff1, s1[i])
            diff2 = append(diff2, s2[i])
        }
    }

    if len(diff1) == 0 {
        return true
    }

    if len(diff1) == 2 {
        return diff1[0] == diff2[1] && diff1[1] == diff2[0]
    }
    return false
}
