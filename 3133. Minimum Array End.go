func minEnd(n int, x int) int64 {
    // insert number n - 1 inside the zero bits of x
    result := int64(x)

    nextNMinus1BitPos := 0
    for i := 0; (int64(1) << nextNMinus1BitPos) <= int64(n - 1); i++ {
        if (result >> i) & 1 == 0 {
            if ((n - 1) >> nextNMinus1BitPos) & 1 == 1 {
                result |= int64(1) << i
            }
            nextNMinus1BitPos++
        }
    }

    return result
}
