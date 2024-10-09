func plusOne(digits []int) []int {
    carry := 1
    for i := len(digits) - 1; i >= 0; i-- {
        newCarry := (digits[i] + carry) / 10
        digits[i] = (digits[i] + carry) % 10
        carry = newCarry
    }

    if carry > 0 {
        digits = prepend(digits, carry)
    }
    return digits
}

func prepend(v []int, val int) []int {
    v = append(v, 0)
    copy(v[1:], v)
    v[0] = val
    return v
}
