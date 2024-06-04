func longestPalindrome(s string) int {
    freq := map[rune]int{}

    for _, c := range(s) {
        freq[c]++
    }

    result := 0
    hasOdd := 0
    for _, qt := range(freq) {
        result += 2 * (qt / 2)
        if qt % 2 == 1 {
            hasOdd = 1
        }
    }
    return result + hasOdd
}
