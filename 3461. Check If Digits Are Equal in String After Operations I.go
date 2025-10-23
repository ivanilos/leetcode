func hasSameDigits(s string) bool {
    v := []int{}
    for _, c := range s {
        v = append(v, int(c - '0'))
    }

    for len(v) > 2 {
        newV := []int{}

        for i := 1; i < len(v); i++ {
            newV = append(newV, (v[i - 1] + v[i]) % 10)
        }
        v = newV
    }
    return v[0] == v[1]
}
