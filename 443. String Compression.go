func compress(chars []byte) int {
    nextPos := 0
    lastChar := chars[0]
    curQt := 0
    for _, c := range chars {
        if c == lastChar {
            curQt++
        } else {
            update(chars, lastChar, curQt, &nextPos)

            lastChar = c
            curQt = 1
        }
    }

    update(chars, lastChar, curQt, &nextPos)

    return nextPos
}

func update(chars []byte, lastChar byte, curQt int, nextPos *int) {
    chars[*nextPos] = lastChar
    *nextPos++
    if curQt > 1 {
        quantityBytes := getQuantityBytes(curQt)

        for _, digit := range quantityBytes {
            chars[*nextPos] = digit
            *nextPos++
        }
    }
}

func getQuantityBytes(curQt int) []byte {
    result := []byte{}
    for curQt > 0 {
        byteVal:= byte((curQt % 10) + int('0'))
        curQt /= 10
        result = append(result, byteVal)
    }

    slices.Reverse(result)
    return result
}
