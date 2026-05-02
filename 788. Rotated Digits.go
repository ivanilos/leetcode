func rotatedDigits(n int) int {
    result := 0
    for i := 1; i <= n; i++ {
        if isGood(i) {
            result++
        }
    }

    return result
}

func isGood(num int) bool {
    freq := map[int]int{}
    for num > 0 {
        digit := num % 10
        num /= 10

        freq[digit]++

        if digit == 3 || digit == 4 || digit == 7 {
            return false
        }
    }

    return freq[2] + freq[5] + freq[6] + freq[9] > 0
}
