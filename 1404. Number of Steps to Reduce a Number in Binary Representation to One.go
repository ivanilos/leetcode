func numSteps(s string) int {
    carry := 0

    result := 0
    for i := len(s) - 1; i >= 0; i-- {
        val := int(s[i]) - int('0') + carry

        if val % 2 == 0 {
            result++
            if s[i] == '1' {
                carry = 1 
            } else {
                carry = 0
            }
        } else {
            if i != 0 {
                result += 2
                carry = 1
            }
        }
    }
    return result
}
