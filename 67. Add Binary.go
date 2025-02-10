func addBinary(a string, b string) string {
    result := []byte{}

    carry := 0
    for i, j := len(a) - 1, len(b) - 1; i >= 0 || j >= 0; i, j = i - 1, j - 1 {
        curSum := 0
        if i >= 0 && a[i] == '1' {
            curSum++
        }
        if j >= 0 && b[j] == '1' {
            curSum++
        }

        nextBit := (curSum + carry) % 2
        result = append(result, byte(nextBit + int('0')))

        carry = (curSum + carry) / 2
    }

    for carry > 0 {
        nextBit := carry % 2
        result = append(result, byte(nextBit + int('0')))
        carry = carry / 2
    }

    slices.Reverse(result)
    return string(result)
}
