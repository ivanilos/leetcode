func minimizeXor(num1 int, num2 int) int {
    ops := getSetBitsQt(num2)

    result := 0

    for i := 31; i >= 0; i-- {
        if ((num1 >> i) & 1) > 0 && ops > 0 {
            result |= 1 << i
            ops--
        }
    }

    for i := 0; i < 32; i++ {
        if ((result >> i) & 1) == 0 && ops > 0 {
            result |= 1 << i
            ops--
        }
    }

    return result

}

func getSetBitsQt(x int) int {
    result := 0
    for i := 0; i < 32; i++ {
        result += (x >> i) & 1
    }
    return result
}
