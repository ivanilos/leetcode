func minimumLength(s string) int {
    freq := map[rune]int{}

    for _, c := range s {
        freq[c]++
    }

    result := 0
    for _, val := range freq {
        if val % 2 == 0 {
            result += 2
        } else {
            result++
        }
    }

    return result
}
