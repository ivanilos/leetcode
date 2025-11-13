func maxOperations(s string) int {
    seenOnes := 0
    zeroOnTheWay := false

    result := 0
    for _, c := range s {
        if c == '1' {
            if zeroOnTheWay {
                result += seenOnes
            }
            seenOnes++
            zeroOnTheWay = false
        } else {
            zeroOnTheWay = true
        }
    }
    if zeroOnTheWay {
        result += seenOnes
    }

    return result
}
