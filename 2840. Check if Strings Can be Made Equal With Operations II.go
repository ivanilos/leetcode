func checkStrings(s1 string, s2 string) bool {
    even1, odd1 := makeCanonical(s1)
    even2, odd2 := makeCanonical(s2)

    return string(even1) == string(even2) && string(odd1) == string(odd2) 
}

func makeCanonical(s1 string) (even []byte, odd []byte) {
    for i := 0; i < len(s1); i++ {
        if i % 2 == 0 {
            even = append(even, s1[i])
        } else {
            odd = append(odd, s1[i])
        }
    }

    slices.Sort(even)
    slices.Sort(odd)
    return even, odd
}
