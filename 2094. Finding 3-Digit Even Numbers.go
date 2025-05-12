func findEvenNumbers(digits []int) []int {
    freq := map[int]int{}

    for _, dig := range digits {
        freq[dig]++
    }

    result := []int{}
    for i := 100; i <= 999; i += 2 {
        if isValid(i, freq) {
            result = append(result, i)
        }
    }

    return result
}

func isValid(num int, freq map[int]int) bool {
    if num % 2 != 0 {
        return false
    }

    needFreq := map[int]int{}

    for num > 0 {
        needFreq[num % 10]++
        num /= 10
    }

    for key, val := range needFreq {
        if freq[key] < val {
            return false
        }
    }

    return true
}
