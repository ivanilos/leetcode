func isOneBitCharacter(bits []int) bool {
    curChar := 0
    lastStart := 0
    for curChar < len(bits) {
        lastStart = curChar
        if bits[curChar] == 1 {
            curChar += 2
        } else {
            curChar++
        }
    }

    if lastStart == len(bits) - 1 {
        return true
    }
    return false
}
