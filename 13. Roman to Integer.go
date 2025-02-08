func romanToInt(s string) int {
    romanToDecimal := map[string]int{
        "I": 1,
        "IV": 4,
        "V": 5,
        "IX": 9,
        "X": 10,
        "XL": 40,
        "L": 50,
        "XC": 90,
        "C": 100,
        "CD": 400,
        "D": 500,
        "CM": 900,
        "M": 1000,
    }

    next := 0
    result := 0
    for next < len(s) {
        fmt.Println(next, result)
        if  hasTwoDigitValue(next, s, romanToDecimal) {
            result += romanToDecimal[s[next : next + 2]]
            next += 2
        } else {
            result += romanToDecimal[s[next : next + 1]]
            next++
        }
    }
    return result
}

func hasTwoDigitValue(next int, s string, romanToDecimal map[string]int) bool {
    if  next + 1 < len(s) {
        _, ok := romanToDecimal[s[next : next + 2]]
        return ok
    }
    return false
}
